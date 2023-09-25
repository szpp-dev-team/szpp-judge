package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/storage"
	"github.com/go-sql-driver/mysql"
	"github.com/szpp-dev-team/szpp-judge/backend/api"
	"github.com/szpp-dev-team/szpp-judge/backend/api/grpc_server"
	"github.com/szpp-dev-team/szpp-judge/backend/core/config"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
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
	testcasesRepository := testcases.NewRepository(storageClient)
	conn, err := grpc.Dial(config.JudgeAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	judgeClient := judgev1.NewJudgeServiceClient(conn)

	// logger
	logger := slog.Default()

	grpcSrv := grpc_server.New(
		api.WithLogger(logger),
		api.WithEntClient(entClient),
		api.WithReflection(config.ModeDev),
		api.WithCloudtasksClient(cloudtasksClient),
		api.WithTestcasesRepository(testcasesRepository),
		api.WithJudgeClient(judgeClient),
	)
	lsnr, err := net.Listen("tcp", "0.0.0.0:"+config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()
	go func() {
		logger.Info("server launched", slog.String("port", config.GrpcPort))
		if err := grpcSrv.Serve(lsnr); err != nil {
			log.Fatal(err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	logger.Info("servers are being stopped")
	grpcSrv.GracefulStop()
}
