package sandboxtest

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"
	"testing"

	"github.com/docker/docker/client"
	"github.com/szpp-dev-team/szpp-judge/judge/sandbox"
	"github.com/szpp-dev-team/szpp-judge/judge/util/fsutil"
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

var (
	dockerClient *client.Client
	sbGcc        *sandbox.Sandbox
	sbGccDev     *sandbox.Sandbox
	sbPython     *sandbox.Sandbox
	sbPythonDev  *sandbox.Sandbox
)

const (
	containerWorkingDir = "/work"
)

func TestMain(m *testing.M) {
	// GitHub Actions などの CI では docker が使えないのでテストをスキップ
	if os.Getenv("CI") == "" {
		os.Exit(0)
	}

	code, err := testMainSub(m)
	if err != nil {
		log.Println("Test failed:", err)
		if code == 0 {
			code = 255
		}
	}
	os.Exit(code)
}

// defer を使うためのサブ関数
func testMainSub(m *testing.M) (int, error) {
	defer func() {
		if sbPythonDev != nil {
			sbPythonDev.Close()
		}
		if sbPython != nil {
			sbPython.Close()
		}
		if sbGccDev != nil {
			sbGccDev.Close()
		}
		if sbGcc != nil {
			sbGcc.Close()
		}
		if dockerClient != nil {
			dockerClient.Close()
		}
	}()

	ctx := context.Background()
	if err := initSandbox(ctx); err != nil {
		return 255, err
	}

	code := m.Run()
	return code, nil
}

func initSandbox(ctx context.Context) error {
	dc, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation(), client.FromEnv)
	if err != nil {
		return err
	}
	dockerClient = dc

	hostBindDir := path.Join(fsutil.GetGoModAbsDir(), "_workdir", "sandboxtest")

	slog.Info("initializing sandboxes",
		slog.String("hostBindDir", hostBindDir),
		slog.String("containerWorkingDir", containerWorkingDir),
		slog.Int("containerProcNumLimit", ContainerProcNumLimit),
		slog.Int64("containerMemoryLimit[MiB]", int64(ContainerMemoryLimit/unit.MiB)),
	)

	if err := os.RemoveAll(hostBindDir); err != nil {
		return fmt.Errorf("Cannot remove %q: %w", hostBindDir, err)
	}
	if err := os.MkdirAll(hostBindDir, 0755); err != nil {
		return fmt.Errorf("Cannot create dir %q: %w", hostBindDir, err)
	}

	setupWorkingDir := sandbox.WithWorkingDir(containerWorkingDir)
	setupBindDir := sandbox.WithBindDir(hostBindDir, containerWorkingDir)
	setupMemoryLimit := sandbox.WithContainerMemoryLimit(ContainerMemoryLimit)
	setupProcNumLimit := sandbox.WithProcNumLimit(ContainerProcNumLimit)
	disableDevMode := sandbox.WithDevMode(false)

	sbGcc, err = sandbox.New(ctx, dc, "szpp-judge-image-gcc13.2",
		setupWorkingDir,
		setupBindDir,
		setupMemoryLimit,
		setupProcNumLimit,
		disableDevMode,
	)
	if err != nil {
		return err
	}
	slog.Info("Created GCC container", slog.String("ID", sbGcc.ContainerID[:6]))

	sbGccDev, err = sandbox.New(ctx, dc, "szpp-judge-image-gcc13.2",
		setupWorkingDir,
		setupBindDir,
		setupMemoryLimit,
		setupProcNumLimit,
		sandbox.WithDevMode(true),
	)
	if err != nil {
		return err
	}
	slog.Info("Created GCC container (devMode)", slog.String("ID", sbGccDev.ContainerID[:6]))

	sbPython, err = sandbox.New(ctx, dc, "szpp-judge-image-cpython3.11",
		setupWorkingDir,
		setupBindDir,
		setupMemoryLimit,
		setupProcNumLimit,
		disableDevMode,
	)
	if err != nil {
		return err
	}
	slog.Info("Created Python container", slog.String("ID", sbPython.ContainerID[:6]))

	sbPythonDev, err = sandbox.New(ctx, dc, "szpp-judge-image-cpython3.11",
		setupWorkingDir,
		setupBindDir,
		setupMemoryLimit,
		setupProcNumLimit,
		sandbox.WithDevMode(true),
	)
	if err != nil {
		return err
	}
	slog.Info("Created Python container (devMode)", slog.String("ID", sbPythonDev.ContainerID[:6]))

	return nil
}
