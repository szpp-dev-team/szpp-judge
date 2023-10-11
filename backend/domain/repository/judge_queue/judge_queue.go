package judge_queue

import (
	"context"
	"fmt"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	judgev1 "github.com/szpp-dev-team/szpp-judge/proto-gen/go/judge/v1"
	"google.golang.org/protobuf/proto"
)

type JudgeQueue interface {
	Push(ctx context.Context, req *judgev1.JudgeRequest) error
}

type cloudtasksImpl struct {
	client              *cloudtasks.Client
	handleJudgeTaskURL  string
	projectID           string
	locationID          string
	queueID             string
	serviceAccountEmail string
}

func New(client *cloudtasks.Client, handleJudgeTaskURL, projectID, locationID, queueID, serviceAccountEmail string) JudgeQueue {
	return &cloudtasksImpl{client, handleJudgeTaskURL, projectID, locationID, queueID, serviceAccountEmail}
}

func (i *cloudtasksImpl) Push(ctx context.Context, req *judgev1.JudgeRequest) error {
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	if _, err := i.client.CreateTask(ctx, &cloudtaskspb.CreateTaskRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s/queues/%s", i.projectID, i.locationID, i.queueID),
		Task: &cloudtaskspb.Task{
			MessageType: &cloudtaskspb.Task_HttpRequest{
				HttpRequest: &cloudtaskspb.HttpRequest{
					Url:        i.handleJudgeTaskURL,
					HttpMethod: cloudtaskspb.HttpMethod_POST,
					Body:       reqBytes,
					AuthorizationHeader: &cloudtaskspb.HttpRequest_OidcToken{
						OidcToken: &cloudtaskspb.OidcToken{
							ServiceAccountEmail: i.serviceAccountEmail,
						},
					},
				},
			},
		},
	}); err != nil {
		return err
	}
	return nil
}
