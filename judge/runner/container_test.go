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
		log.Fatal(err)
	}

	r, err := NewDockerContaineredRunner(ctx, dc, lang.GCC, cfg.WorkingDirAbsRoot)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	log.Printf("Successfully created a container: ID=%s, hostWorkingDir=%s\n", r.ID, r.HostWorkingDir)
}
