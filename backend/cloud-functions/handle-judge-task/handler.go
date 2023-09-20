package handlejudgetask

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/go-sql-driver/mysql"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	"github.com/szpp-dev-team/szpp-judge/backend/usecases/judge"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func init() {
	config, err := newConfig()
	if err != nil {
		log.Fatal(err)
	}
	judgeConn, err := grpc.Dial(config.JudgeAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer judgeConn.Close()
	judgeClient := judgev1.NewJudgeServiceClient(judgeConn)
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
	interactor := judge.NewInteractor(judgeClient, entClient)
	functions.HTTP("HandleJudgeTask", HandleJudgeTask(interactor))
}

func HandleJudgeTask(interactor *judge.Interactor) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var req judgev1.JudgeRequest
		if err := proto.Unmarshal(b, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := interactor.PostJudgeRequest(r.Context(), &req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
