package judge

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	docker "github.com/docker/docker/client"
	"github.com/google/uuid"
	"github.com/samber/lo"
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

	langMeta, ok := langs.Get(langs.LangID(req.LangId))
	if !ok {
		return status.Error(codes.InvalidArgument, "no such language")
	}

	workdir := filepath.Join(i.workdirRoot, uuid.NewString())
	if err := os.MkdirAll(workdir, 0755); err != nil {
		i.logger.Error("failed to create workdir", slog.Any("error", err))
		return status.Error(codes.Internal, err.Error())
	}
	defer os.RemoveAll(workdir)

	sb, err := sandbox.New(ctx, i.dockerClient, langMeta.DockerImage,
		sandbox.WithWorkingDir("/work"),
		sandbox.WithBindDir(workdir, "/work"),
		sandbox.WithContainerMemoryLimit(unit.Byte(req.ExecMemoryLimitMib)*unit.MiB+unit.GiB),
	)
	if err != nil {
		i.logger.Error("failed to instance sandbox", slog.Any("error", err))
		return status.Error(codes.Internal, err.Error())
	}
	defer sb.Close()

	if err := i.downloadSourceCodeToFile(ctx, filepath.Join(workdir, langMeta.SourceFile), req.SourceCodePath); err != nil {
		return err
	}

	ok, compileMessage, err := i.compile(ctx, langMeta, sb)
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

	for _, tc := range req.Testcases {
		testcase, err := i.downloadTestcase(ctx, tc.InputPath, tc.ExpectedPath)
		if err != nil {
			return err
		}

		res, err := sb.Exec(ctx, sandbox.ExecOption{
			Stdin: bytes.NewBuffer(testcase.Input),
			Cmd:   langMeta.ExecCmd,
		}, sandbox.SzpprunOption{
			TimeLimit:   time.Duration(req.ExecTimeLimitMs+200) * time.Millisecond,
			MemoryLimit: unit.Byte(req.ExecMemoryLimitMib) * unit.MiB,
		})
		if err != nil {
			i.logger.Error("error occurred while exec", slog.Any("error", err))
			return status.Error(codes.Internal, err.Error())
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
			i.logger.Error("RE occurred", slog.Any("exitCode", res.ExitCode), slog.Any("stderr", res.Stderr))
			judgeResp.Status = judgev1.JudgeStatus_RE
		case res.ExecTime > time.Duration(req.ExecTimeLimitMs)*time.Millisecond:
			judgeResp.Status = judgev1.JudgeStatus_TLE
		case res.ExecMemory > unit.Byte(req.ExecMemoryLimitMib):
			judgeResp.Status = judgev1.JudgeStatus_MLE
		case res.StdoutOverflowed:
			judgeResp.Status = judgev1.JudgeStatus_OLE
		default:
			checker := &Checker{
				Output:       res.Stdout,
				ExpectOutput: string(testcase.Output),
			}
			switch ty := req.JudgeType.JudgeType.(type) {
			case *judgev1.JudgeType_Normal:
				{
					if checker.JudgeNormal(lo.FromPtrOr(ty.Normal.CaseInsensitive, false)) {
						judgeResp.Status = judgev1.JudgeStatus_AC
					} else {
						judgeResp.Status = judgev1.JudgeStatus_WA
					}
				}
			case *judgev1.JudgeType_Eps, *judgev1.JudgeType_Custom, *judgev1.JudgeType_Interactive:
				i.logger.Error("unsupported judge type", slog.Any("type", ty))
				return status.Error(codes.Unimplemented, "unsupported judge type")
			}
		}

		if err := stream.Send(judgeResp); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
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
		Cmd: langMeta.CompileCmd,
	}, sandbox.SzpprunOption{
		TimeLimit:      time.Minute,
		FileWriteLimit: unit.MiB,
	})
	if err != nil {
		i.logger.Error("error occurred while compile", slog.Any("error", err))
		return false, "", status.Error(codes.Internal, err.Error())
	}
	return res.ExitCode == 0, res.Stderr, nil
}

func (i *Interactor) downloadSourceCodeToFile(ctx context.Context, filename, gcsPath string) error {
	f, err := os.Create(filename)
	if err != nil {
		i.logger.Error("failed to create source code file", slog.Any("error", err))
		return status.Error(codes.Internal, err.Error())
	}
	defer f.Close()
	r, err := i.gcsClient.Bucket("szpp-judge").Object(gcsPath).NewReader(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
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
	inputReader, err := i.gcsClient.Bucket("szpp-judge").Object(inputPath).NewReader(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		i.logger.Error("failed to open input", slog.Any("error", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer inputReader.Close()
	outputReader, err := i.gcsClient.Bucket("szpp-judge").Object(outputPath).NewReader(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
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
