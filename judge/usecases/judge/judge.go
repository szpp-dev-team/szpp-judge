package judge

import (
	"bytes"
	"context"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	docker "github.com/docker/docker/client"
	"github.com/google/uuid"
	"github.com/szpp-dev-team/szpp-judge/judge/sandbox"
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
	"github.com/szpp-dev-team/szpp-judge/langs"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Interactor struct {
	dockerClient *docker.Client
	gcsClient    *storage.Client
	workdirRoot  string
	bucketName   string
	logger       *slog.Logger
}

func NewInteractor(dockerClient *docker.Client, gcsClient *storage.Client, workdirRoot, bucketName string) *Interactor {
	if !filepath.IsAbs(workdirRoot) {
		panic("workdir must be absolute path")
	}
	logger := slog.Default()
	return &Interactor{dockerClient, gcsClient, workdirRoot, bucketName, logger}
}

func (i *Interactor) Judge(req *judgev1.JudgeRequest, stream judgev1.JudgeService_JudgeServer) error {
	if err := i.judgeMain(req, stream); err != nil {
		if err := sendIE(stream, req.SubmissionId); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (i *Interactor) judgeMain(req *judgev1.JudgeRequest, stream judgev1.JudgeService_JudgeServer) error {
	ctx := stream.Context()

	lm, ok := langs.Get(langs.LangID(req.LangId))
	if !ok {
		return status.Error(codes.InvalidArgument, "no such language")
	}

	workdir := filepath.Join(i.workdirRoot, uuid.NewString())
	if err := os.MkdirAll(workdir, 0666); err != nil {
		i.logger.Error("failed to create workdir", slog.Any("error", err))
		return status.Error(codes.Internal, err.Error())
	}
	defer os.RemoveAll(workdir)

	sb, err := sandbox.New(ctx, i.dockerClient, lm.DockerImage,
		sandbox.WithWorkingDir("/work"),
		sandbox.WithBindDir(workdir, "/work"),
		sandbox.WithContainerMemoryLimit(unit.Byte(req.ExecMemoryLimitMib)*unit.MiB+unit.GiB),
	)
	if err != nil {
		i.logger.Error("failed to instance sandbox", slog.Any("error", err))
		return status.Error(codes.Internal, err.Error())
	}
	defer sb.Close()

	// download and compile source code
	ok, compileMessage, err := i.downloadAndCompileCode(ctx, filepath.Join(workdir, lm.SourceFile), req.SourceCodePath, lm, sb)
	if err != nil {
		return err
	}
	if !ok {
		if err := stream.Send(&judgev1.JudgeResponse{
			SubmissionId:    req.SubmissionId,
			Status:          judgev1.JudgeStatus_CE,
			CompilerMessage: compileMessage,
		}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		return nil
	}

	// download and compile checker
	checkerLm, _ := langs.Get(langs.CPP_20_GCC_TESTLIB)
	ok, _, err = i.downloadAndCompileCode(ctx, filepath.Join(workdir, "checker.cpp"), req.CheckerCodePath, checkerLm, sb)
	if err != nil {
		return err
	}
	if !ok {
		return status.Error(codes.Internal, "failed to compile checker")
	}

	for _, tc := range req.Testcases {
		testcase, err := i.downloadTestcase(ctx, tc.InputPath, tc.ExpectedPath)
		if err != nil {
			return err
		}

		res, err := i.run(ctx, workdir, req, testcase, lm, sb)
		if err != nil {
			return err
		}

		judgeResp := &judgev1.JudgeResponse{
			SubmissionId:  req.SubmissionId,
			Status:        judgev1.JudgeStatus_AC,
			TestcaseId:    tc.Id,
			ExecTimeMs:    uint32(res.ExecTime.Milliseconds()),
			ExecMemoryKib: uint32(res.ExecMemory / unit.KiB),
		}
		switch {
		case res.ExitCode != 0:
			i.logger.Warn("RE occurred", slog.Any("exitCode", res.ExitCode), slog.Any("stderr", res.Stderr))
			judgeResp.Status = judgev1.JudgeStatus_RE
		case res.ExecTime > time.Duration(req.ExecTimeLimitMs)*time.Millisecond:
			judgeResp.Status = judgev1.JudgeStatus_TLE
		case res.ExecMemory > unit.Byte(req.ExecMemoryLimitMib)*unit.MiB:
			log.Println(res.ExecMemory, unit.Byte(req.ExecMemoryLimitMib))
			judgeResp.Status = judgev1.JudgeStatus_MLE
		case res.StdoutOverflowed:
			judgeResp.Status = judgev1.JudgeStatus_OLE
		default:
			ok, err := i.check(ctx, checkerLm, sb)
			if err != nil {
				return err
			}
			if !ok {
				judgeResp.Status = judgev1.JudgeStatus_WA
			}
		}

		if err := stream.Send(judgeResp); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (i *Interactor) run(
	ctx context.Context,
	workdir string,
	req *judgev1.JudgeRequest,
	testcase *Testcase,
	lm *langs.Meta,
	sb *sandbox.Sandbox,
) (*sandbox.ExecResult, error) {
	res, err := sb.Exec(ctx, sandbox.ExecOption{
		AsRootUser:      false,
		Stdin:           bytes.NewBuffer(testcase.Input),
		Cmd:             lm.ExecCmd,
		StdoutReadLimit: 256 * unit.KiB,
		StderrReadLimit: 64 * unit.KiB,
	}, sandbox.SzpprunOption{
		TimeLimit:        time.Duration(req.ExecTimeLimitMs+200) * time.Millisecond,
		MemoryLimit:      unit.Byte(req.ExecMemoryLimitMib) * unit.MiB,
		FileWriteLimit:   unit.MiB * 8,
		NumOpenFileLimit: 128,
	})
	if err != nil {
		i.logger.Error("error occurred while exec", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := saveAsFile(filepath.Join(workdir, "testcase_input.txt"), testcase.Input); err != nil {
		return nil, err
	}
	if err := saveAsFile(filepath.Join(workdir, "testcase_output.txt"), testcase.Output); err != nil {
		return nil, err
	}
	if err := saveAsFile(filepath.Join(workdir, "user_output.txt"), []byte(res.Stdout)); err != nil {
		return nil, err
	}

	return &res, nil
}

func saveAsFile(name string, b []byte) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, bytes.NewReader(b))
	return err
}

func sendIE(stream judgev1.JudgeService_JudgeServer, submissionID int32) error {
	if err := stream.Send(&judgev1.JudgeResponse{
		SubmissionId: submissionID,
		Status:       judgev1.JudgeStatus_IE,
	}); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (i *Interactor) compile(ctx context.Context, langMeta *langs.Meta, sb *sandbox.Sandbox) (bool, string, error) {
	res, err := sb.Exec(ctx, sandbox.ExecOption{
		AsRootUser:          true,
		Stdin:               nil,
		Cmd:                 langMeta.CompileCmd,
		StdoutReadLimit:     16 * unit.KiB,
		StderrReadLimit:     unit.KiB,
		MergeStderrToStdout: true,
	}, sandbox.SzpprunOption{
		TimeLimit:        time.Minute,
		MemoryLimit:      512 * unit.MiB,
		FileWriteLimit:   1024 * unit.MiB,
		NumOpenFileLimit: 1024,
	})
	if err != nil {
		i.logger.Error("error occurred while compile", slog.Any("error", err))
		return false, "", status.Error(codes.Internal, err.Error())
	}
	i.logger.Info("compile finished", slog.Any("result", res))
	return res.ExitCode == 0, res.Stderr, nil
}

func (i *Interactor) downloadAndCompileCode(ctx context.Context, dst, gcsPath string, lm *langs.Meta, sb *sandbox.Sandbox) (bool, string, error) {
	if err := i.downloadAsFile(ctx, dst, gcsPath); err != nil {
		return false, "", err
	}
	return i.compile(ctx, lm, sb)
}

func (i *Interactor) downloadAsFile(ctx context.Context, filename, gcsPath string) error {
	f, err := os.Create(filename)
	if err != nil {
		i.logger.Error("failed to create source code file", slog.Any("error", err))
		return status.Error(codes.Internal, err.Error())
	}
	defer f.Close()
	r, err := i.gcsClient.Bucket(i.bucketName).Object(gcsPath).NewReader(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			i.logger.Warn("object does not exist", slog.String("bucketName", i.bucketName), slog.String("path", gcsPath))
			return status.Error(codes.NotFound, err.Error())
		}
		i.logger.Error("failed to open source code", slog.Any("error", err))
		return status.Error(codes.Internal, err.Error())
	}
	defer r.Close()
	if _, err := io.Copy(f, r); err != nil {
		i.logger.Error("failed to read source code", slog.Any("error", err))
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

type Testcase struct {
	Input  []byte
	Output []byte
}

func (i *Interactor) downloadTestcase(ctx context.Context, inputPath, outputPath string) (*Testcase, error) {
	inputReader, err := i.gcsClient.Bucket(i.bucketName).Object(inputPath).NewReader(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			i.logger.Warn("object does not exist", slog.String("bucketName", i.bucketName), slog.String("path", inputPath))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		i.logger.Error("failed to open input", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer inputReader.Close()
	outputReader, err := i.gcsClient.Bucket(i.bucketName).Object(outputPath).NewReader(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			i.logger.Warn("object does not exist", slog.String("bucketName", i.bucketName), slog.String("path", outputPath))
			return nil, status.Error(codes.NotFound, err.Error())
		}
		i.logger.Error("failed to open output", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer outputReader.Close()
	input, err := io.ReadAll(inputReader)
	if err != nil {
		i.logger.Error("failed to read input", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	output, err := io.ReadAll(outputReader)
	if err != nil {
		i.logger.Error("failed to read output", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &Testcase{Input: input, Output: output}, nil
}

func (i *Interactor) check(ctx context.Context, lm *langs.Meta, sb *sandbox.Sandbox) (bool, error) {
	res, err := sb.Exec(ctx, sandbox.ExecOption{
		AsRootUser:          true,
		Stdin:               nil,
		Cmd:                 lm.ExecCmd,
		StdoutReadLimit:     16 * unit.KiB,
		MergeStderrToStdout: true,
	}, sandbox.SzpprunOption{
		TimeLimit:        time.Minute,
		MemoryLimit:      512 * unit.MiB,
		FileWriteLimit:   unit.MiB,
		NumOpenFileLimit: 1024,
	})
	if err != nil {
		i.logger.Error("error occurred while compile", slog.Any("error", err))
		return false, status.Error(codes.Internal, err.Error())
	}
	slog.Info("check was finished", slog.Any("result", res))
	return res.ExitCode == 0, nil
}
