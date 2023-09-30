package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/storage"
	docker "github.com/docker/docker/client"
	"github.com/szpp-dev-team/szpp-judge/judge/config"
	"github.com/szpp-dev-team/szpp-judge/judge/servers/grpc_server"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// logger
	logger := slog.Default()

	// Docker Client
	dockerClient, err := docker.NewClientWithOpts(
		docker.FromEnv,
		docker.WithAPIVersionNegotiation(),
	)
	if err != nil {
		log.Fatal(err)
	}

	// GCS Client
	gcsClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc_server.New(
		grpc_server.WithLogger(logger),
		grpc_server.WithDockerClient(dockerClient),
		grpc_server.WithReflection(cfg.ModeDev),
		grpc_server.WithGcsClient(gcsClient),
		grpc_server.WithWorkdirRoot("./_workdir"),
	)

	lsnr, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		gcsClient.Close()
		dockerClient.Close()
		log.Fatal(err)
	}
	defer gcsClient.Close()
	defer dockerClient.Close()
	defer lsnr.Close()
	go func() {
		logger.Info("launched server:", "GrpcPort", cfg.GrpcPort, "ModeDev", cfg.ModeDev)
		if err := srv.Serve(lsnr); err != nil {
			log.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	logger.Info("gracefully stopping the server")
	srv.GracefulStop()
}
