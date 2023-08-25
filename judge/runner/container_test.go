package runner

import (
	"context"
	"log"
	"testing"

	docker "github.com/docker/docker/client"
	"github.com/szpp-dev-team/szpp-judge/judge/config"
	"github.com/szpp-dev-team/szpp-judge/judge/runner/lang"
)

func loadConfig() *config.Config {
	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded config: %#+v", c)
	return c
}

func TestNewDockerContaineredRunner(t *testing.T) {
	ctx := context.Background()
	cfg := loadConfig()

	dc, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		t.Fatal(err)
	}

	r, err := NewDockerContaineredRunner(ctx, dc, lang.GCC, cfg.WorkingDirAbsRoot)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created a container: ID=%s, hostWorkingDir=%s\n", r.ContainerID, r.HostWorkingDir)

	execResult, err := r.Exec(ctx, []string{"curl", "http://example.com"})
	if err != nil {
		r.Close()
		t.Fatal(err)
	}

	t.Log("------------ execution result ------------")
	t.Logf("::: stdout :::\n'%s'\n", execResult.Stdout)
	t.Logf("::: stderr :::\n'%s'\n", execResult.Stderr)
	t.Logf("::: exit code :::\n[%d]\n", execResult.ExitCode)
	t.Log("------------------------------------------")

	defer r.Close()
}
