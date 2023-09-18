package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"github.com/go-sql-driver/mysql"
	"github.com/szpp-dev-team/szpp-judge/backend/api"
	"github.com/szpp-dev-team/szpp-judge/backend/api/grpc_server"
	"github.com/szpp-dev-team/szpp-judge/backend/core/config"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
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
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	// api clients
	cloudtasksClient, err := cloudtasks.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer cloudtasksClient.Close()
	judgeConn, err := grpc.Dial(config.JudgeAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer judgeConn.Close()
	judgeClient := judgev1.NewJudgeServiceClient(judgeConn)

	// logger
	logger := slog.Default()

	srv := grpc_server.New(
		api.WithLogger(logger),
		api.WithEntClient(entClient),
		api.WithReflection(config.ModeDev),
		api.WithJudgeClient(judgeClient),
	)
	lsnr, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()
	go func() {
		logger.Info("server launched")
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
