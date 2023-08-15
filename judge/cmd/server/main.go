package main

import (
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	docker "github.com/docker/docker/client"
	"github.com/szpp-dev-team/szpp-judge/judge/config"
	"github.com/szpp-dev-team/szpp-judge/judge/servers/grpc_server"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// Time zone
	_, err = time.LoadLocation("Asia/Tokyo")
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

	srv := grpc_server.New(
		grpc_server.WithLogger(logger),
		grpc_server.WithDockerClient(dockerClient),
		grpc_server.WithReflection(cfg.ModeDev),
	)
	lsnr, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
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
