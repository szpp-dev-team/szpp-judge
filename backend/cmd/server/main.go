package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/storage"
	"github.com/go-sql-driver/mysql"
	"github.com/szpp-dev-team/szpp-judge/backend/api/connect_server"
	"github.com/szpp-dev-team/szpp-judge/backend/core/config"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/checkers"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/judge_queue"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/sources"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/testcases"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// api clients
	cloudtasksClient, err := cloudtasks.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cloudtasksClient.Close()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer storageClient.Close()
	testcasesRepository := testcases.NewRepository(storageClient, config.GcsBucketName)
	sourcesRepository := sources.NewRepository(storageClient, config.GcsBucketName)
	checkerRepository := checkers.NewRepository(storageClient, config.GcsBucketName)
	judgeQueue := judge_queue.New(
		cloudtasksClient,
		config.HandleJudgeTaskURL,
		config.CloudTasksProjectID,
		config.CloudTasksLocationID,
		config.CloudTasksQueueID,
		config.ServiceAccountEmail,
	)
	conn, err := grpc.Dial(config.JudgeAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	judgeClient := judgev1.NewJudgeServiceClient(conn)

	// logger
	logger := slog.Default()

	srv := connect_server.New(
		fmt.Sprintf("0.0.0.0:%s", config.ConnectPort),
		connect_server.WithLogger(logger),
		connect_server.WithEntClient(entClient),
		connect_server.WithSourcesRepository(sourcesRepository),
		connect_server.WithTestcasesRepository(testcasesRepository),
		connect_server.WithJudgeQueue(judgeQueue),
		connect_server.WithJudgeClient(judgeClient),
		connect_server.WithCheckerRepository(checkerRepository),
		connect_server.WithSecret(config.JWTSecret),
		connect_server.WithFrontendURL(config.FrontendURL),
	)

	go func() {
		logger.Info("server launched", slog.String("port", config.ConnectPort))
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	logger.Info("server is being stopped")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}
