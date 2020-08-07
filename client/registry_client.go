package client

import (
	"context"
	"github.com/oasisprotocol/oasis-core/go/common/node"
	"github.com/oasisprotocol/oasis-core/go/registry/api"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	"google.golang.org/grpc"
	"time"
)

var (
	_ RegistryClient = (*registryClient)(nil)
)

type RegistryClient interface {
	GetNodeById(context.Context, signature.PublicKey, int64) (*node.Node, error)
}

func NewRegistryClient(conn *grpc.ClientConn) *registryClient {
	return &registryClient{
		client: api.NewRegistryClient(conn),
	}
}

type registryClient struct {
	client api.Backend
}

func (c *registryClient) GetNodeById(ctx context.Context, key signature.PublicKey, height int64) (*node.Node, error) {
	defer logRequestDuration(time.Now(), "RegistryClient_GetNodeById")

	q := &api.IDQuery{
		Height: height,
		ID:     key,
	}
	return c.client.GetNode(ctx, q)
}
