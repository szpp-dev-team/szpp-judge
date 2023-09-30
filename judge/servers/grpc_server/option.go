package grpc_server

import (
	"log/slog"

	"cloud.google.com/go/storage"
	docker "github.com/docker/docker/client"
)

type option struct {
	logger        *slog.Logger
	dockerClient  *docker.Client
	gcsClient     *storage.Client
	workdirRoot   string
	useReflection bool
}

func defaultOption() *option {
	return &option{
		logger: slog.Default(),
	}
}

type optionFunc func(*option)

func WithLogger(logger *slog.Logger) optionFunc {
	return func(o *option) {
		o.logger = logger
	}
}

func WithDockerClient(c *docker.Client) optionFunc {
	return func(o *option) {
		o.dockerClient = c
	}
}

// publish grpc server information(method, service, etc.)
func WithReflection(b bool) optionFunc {
	return func(o *option) {
		o.useReflection = b
	}
}

func WithGcsClient(c *storage.Client) optionFunc {
	return func(o *option) {
		o.gcsClient = c
	}
}

func WithWorkdirRoot(s string) optionFunc {
	return func(o *option) {
		o.workdirRoot = s
	}
}
