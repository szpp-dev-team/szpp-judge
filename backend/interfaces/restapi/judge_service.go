package restapi

import (
	"io"
	"net/http"

	"github.com/szpp-dev-team/szpp-judge/backend/usecases/judge"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/protobuf/proto"
)

type JudgeService struct {
	interactor *judge.Interactor
}

func NewJudgeService(interactor *judge.Interactor) *JudgeService {
	return &JudgeService{interactor}
}

func (s *JudgeService) PostJudgeRequest(w http.ResponseWriter, r *http.Request) {
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

	if err := s.interactor.PostJudgeRequest(r.Context(), &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
