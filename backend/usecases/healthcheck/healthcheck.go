package healthcheck

import (
	pb "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

type Interactor struct{}

func NewInteractor() *Interactor {
	return &Interactor{}
}

func (i *Interactor) Ping(req *pb.PingRequest) *pb.PingResponse {
	return &pb.PingResponse{
		Message: req.Name + " pong:)",
	}
}
