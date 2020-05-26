package grpc

import (
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountserver"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockserver"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/chain/chainpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/chain/chainserver"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/stateserver"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionserver"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorserver"
	"github.com/figment-networks/oasis-rpc-proxy/utils/logger"
	"github.com/oasislabs/oasis-core/go/genesis/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GRPC struct {
	server *grpc.Server

	doc    *api.Document
	client *client.Client
}

func New(doc *api.Document, c *client.Client) *GRPC {
	app := &GRPC{
		server: grpc.NewServer(),
		doc: doc,
		client: c,
	}
	return app.init()
}

func (g *GRPC) init() *GRPC {
	logger.Info("initializing grpc servers...")

	chainpb.RegisterChainServiceServer(g.server, chainserver.New(g.doc))
	accountpb.RegisterAccountServiceServer(g.server, accountserver.New(g.client))
	blockpb.RegisterBlockServiceServer(g.server, blockserver.New(g.client))
	statepb.RegisterStateServiceServer(g.server, stateserver.New(g.client))
	validatorpb.RegisterValidatorServiceServer(g.server, validatorserver.New(g.client))
	transactionpb.RegisterTransactionServiceServer(g.server, transactionserver.New(g.client))

	// Register reflection service on gRPC server.
	reflection.Register(g.server)

	return g
}

func (g *GRPC) Start(serverAddress string, serverPort int64) error {
	logger.Info(fmt.Sprintf("starting grpc server %s:%d...", serverAddress, serverPort), logger.Field("app", "grpc"))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverAddress, serverPort))
	if err != nil {
		return err
	}
	return g.server.Serve(lis)
}
