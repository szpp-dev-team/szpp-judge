package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"github.com/go-sql-driver/mysql"
	"github.com/szpp-dev-team/szpp-judge/backend/api"
	"github.com/szpp-dev-team/szpp-judge/backend/api/grpc_server"
	"github.com/szpp-dev-team/szpp-judge/backend/api/restapi_server"
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

	httpSrv := &http.Server{
		Handler: restapi_server.New(api.WithLogger(logger), api.WithEntClient(entClient), api.WithJudgeClient(judgeClient)),
		Addr:    ":" + config.HttpPort,
	}
	go func() {
		logger.Info("http server launched")
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	grpcSrv := grpc_server.New(
		api.WithLogger(logger),
		api.WithEntClient(entClient),
		api.WithReflection(config.ModeDev),
		api.WithJudgeClient(judgeClient),
		api.WithCloudtasksClient(cloudtasksClient),
	)
	lsnr, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()
	go func() {
		logger.Info("grpc server launched")
		if err := grpcSrv.Serve(lsnr); err != nil {
			log.Fatal(err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	logger.Info("servers are being stopped")
	grpcSrv.GracefulStop()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}
