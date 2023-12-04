package client

import (
	"fmt"
	"github.com/jjangg96/oasis-rpc-proxy/metric"
	"github.com/jjangg96/oasis-rpc-proxy/utils/logger"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	oasisGrpc "github.com/oasisprotocol/oasis-core/go/common/grpc"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
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

func getAddress(rawAddress string) (*api.Address, error) {
	address := api.Address{}
	if err := address.UnmarshalText([]byte(rawAddress)); err != nil {
		return nil, err
	}
	return &address, nil
}

func getAddressFromPublicKey(key string) (*api.Address, error) {
	var pk signature.PublicKey
	if err := pk.UnmarshalText([]byte(key)); err != nil {
		return nil, err
	}
	address := api.NewAddress(pk)
	return &address, nil
}

func logRequestDuration(start time.Time, requestName string) {
	elapsed := time.Since(start)
	metric.ClientRequestDuration.WithLabelValues(requestName).Set(elapsed.Seconds())
}
