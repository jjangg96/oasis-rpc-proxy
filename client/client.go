package client

import (
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/metric"
	"github.com/figment-networks/oasis-rpc-proxy/utils/logger"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	oasisGrpc "github.com/oasislabs/oasis-core/go/common/grpc"
	"google.golang.org/grpc"
	"time"
)

func New(target string) (*Client, error) {
	logger.Debug(fmt.Sprintf("grpc server target is %s", target))

	conn, err := oasisGrpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,

		Consensus: NewConsensusClient(conn),
		Registry:  NewRegistryClient(conn),
		Scheduler: NewSchedulerClient(conn),
		Staking:   NewStakingClient(conn),
	}, nil
}

type Client struct {
	conn *grpc.ClientConn

	Consensus ConsensusClient
	Registry  RegistryClient
	Scheduler SchedulerClient
	Staking   StakingClient
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func getPublicKey(key string) (*signature.PublicKey, error) {
	var pKey signature.PublicKey
	if err := pKey.UnmarshalText([]byte(key)); err != nil {
		return nil, err
	}
	return &pKey, nil
}

func logRequestDuration(start time.Time, requestName string) {
	elapsed := time.Since(start)
	metric.ClientRequestDuration.WithLabelValues(requestName).Set(elapsed.Seconds())
}
