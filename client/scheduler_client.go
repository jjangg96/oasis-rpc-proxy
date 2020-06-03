package client

import (
	"context"
	"github.com/oasislabs/oasis-core/go/scheduler/api"
	"google.golang.org/grpc"
	"time"
)

var (
	_ SchedulerClient = (*schedulerClient)(nil)
)

type SchedulerClient interface {
	GetValidatorsByHeight(context.Context, int64) ([]*api.Validator, error)
}

func NewSchedulerClient(conn *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{
		client:   api.NewSchedulerClient(conn),
	}
}

type schedulerClient struct {
	client   api.Backend
}

func (r *schedulerClient) GetValidatorsByHeight(ctx context.Context, h int64) ([]*api.Validator, error) {
	defer logRequestDuration(time.Now(), "SchedulerClient_GetValidatorsByHeight")

	return r.client.GetValidators(ctx, h)
}
