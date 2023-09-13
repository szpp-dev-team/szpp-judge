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
	Push(ctx context.Context, submitID int, req *judgev1.JudgeRequest) error
}

type cloudtasksImpl struct {
	client    *cloudtasks.Client
	projectID string
	location  string
}

func New(client *cloudtasks.Client, projectID, location string) JudgeQueue {
	return &cloudtasksImpl{client, projectID, location}
}

func (i *cloudtasksImpl) Push(ctx context.Context, submitID int, req *judgev1.JudgeRequest) error {
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	if _, err := i.client.CreateTask(ctx, &cloudtaskspb.CreateTaskRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s/queues/%d", i.projectID, i.location, submitID),
		Task: &cloudtaskspb.Task{
			MessageType: &cloudtaskspb.Task_AppEngineHttpRequest{
				AppEngineHttpRequest: &cloudtaskspb.AppEngineHttpRequest{
					HttpMethod:  cloudtaskspb.HttpMethod_POST,
					RelativeUri: "/judge",
					Body:        reqBytes,
				},
			},
		},
	}); err != nil {
		return err
	}
	return nil
}
