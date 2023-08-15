package grpc_server

import (
	"log/slog"

	docker "github.com/docker/docker/client"
)

type option struct {
	logger        *slog.Logger
	dockerClient  *docker.Client
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
