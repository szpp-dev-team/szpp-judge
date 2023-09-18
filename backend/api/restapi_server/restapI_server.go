package restapi_server

import (
	"net/http"

	"github.com/szpp-dev-team/szpp-judge/backend/api"
	restapi_interfaces "github.com/szpp-dev-team/szpp-judge/backend/interfaces/restapi"
	u_judge "github.com/szpp-dev-team/szpp-judge/backend/usecases/judge"
)

func New(opts ...api.OptionFunc) *http.ServeMux {
	opt := api.DefaultOption()
	for _, f := range opts {
		f(opt)
	}

	judgeInteractor := u_judge.NewInteractor(opt.JudgeClient, opt.EntClient)
	judgeService := restapi_interfaces.NewJudgeService(judgeInteractor)

	mux := http.NewServeMux()
	mux.HandleFunc("/post-judge-request", judgeService.PostJudgeRequest)

	return mux
}
