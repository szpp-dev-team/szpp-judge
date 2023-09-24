package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/storage"
	"github.com/go-sql-driver/mysql"
	"github.com/szpp-dev-team/szpp-judge/backend/api/grpc_server"
	"github.com/szpp-dev-team/szpp-judge/backend/core/config"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	"golang.org/x/exp/slog"
)

func main() {
	ctx := context.Background()

	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// mysql client
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConfig := &mysql.Config{
		DBName:               config.DBName,
		User:                 config.DBUser,
		Passwd:               config.DBPass,
		Addr:                 config.DBAddr,
		Net:                  "tcp",
		ParseTime:            true,
		Loc:                  jst,
		AllowNativePasswords: true,
	}
	entClient, err := ent.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer entClient.Close()
	if err := entClient.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	// logger
	logger := slog.Default()

	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer storageClient.Close()
	testcasesRepository := testcases.NewRepository(storageClient)

	srv := grpc_server.New(
		grpc_server.WithLogger(logger),
		grpc_server.WithEntClient(entClient),
		grpc_server.WithReflection(config.ModeDev),
		grpc_server.WithTestcasesRepository(testcasesRepository),
	)
	lsnr, err := net.Listen("tcp", "0.0.0.0:"+config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()
	go func() {
		logger.Info("server launched", slog.String("port", config.GrpcPort))
		if err := srv.Serve(lsnr); err != nil {
			log.Fatal(err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	logger.Info("server is being stopped")
	srv.GracefulStop()
}
