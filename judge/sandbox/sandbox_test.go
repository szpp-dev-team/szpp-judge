package sandbox

import (
	"context"
	"log"
	"testing"

	docker "github.com/docker/docker/client"
	"github.com/szpp-dev-team/szpp-judge/judge/config"
)

func loadDevConfig() *config.DevConfig {
	c, err := config.NewDevConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded dev config: %#+v", c)
	return c
}

func TestNew_Close(t *testing.T) {
	ctx := context.Background()
	cfg := loadDevConfig()

	dc, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		t.Fatal(err)
	}
	defer dc.Close()

	sb, err := New(ctx, dc, "szpp-judge-image-gcc13.2",
		WithWorkingDir("/work"),
		WithBindDir(cfg.WorkingDirAbsRoot, "/work"),
	)
	if err != nil {
		dc.Close()
		t.Fatal(err)
	}
	t.Logf("Successfully created a container: ID=%s, hostWorkingDir=%s\n", sb.ContainerID, sb.HostBindDir)

	execResult, err := sb.ExecRaw(ctx, ExecOption{
		AsRootUser: false,
		Stdin:      nil,
		Cmd:        []string{"curl", "http://example.com"},
	})
	if err != nil {
		sb.Close()
		t.Fatal(err)
	}
	defer sb.Close()

	t.Log("------------ execution result ------------")
	t.Logf("::: stdout :::\n'%s'\n", execResult.Stdout)
	t.Logf("::: stderr :::\n'%s'\n", execResult.Stderr)
	t.Logf("::: exit code :::\n[%d]\n", execResult.ExitCode)
	t.Log("------------------------------------------")
}
