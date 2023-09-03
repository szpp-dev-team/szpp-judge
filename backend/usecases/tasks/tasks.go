package tasks

import (
	"context"

	"github.com/szpp-dev-team/szpp-judge/backend/core/entutil"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent"
	backendv1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/backend/v1"
)

type Interactor struct {
	entClient *ent.Client
}

func NewInteractor(entClient *ent.Client) *Interactor {
	return &Interactor{entClient}
}

func (i *Interactor) CreateTask(ctx context.Context, req *backendv1.CreateTaskRequest) (*backendv1.CreateTaskResponse, error) {
	entutil.WithTx(ctx, i.entClient, func(tx *ent.Tx) error {
		tx.Client()
	})
}
