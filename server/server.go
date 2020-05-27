package server

import (
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/chain/chainpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/figment-networks/oasis-rpc-proxy/utils/logger"
	"github.com/oasislabs/oasis-core/go/genesis/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	server *grpc.Server

	doc    *api.Document
	client *client.Client
}

func New(doc *api.Document, c *client.Client) *Server {
	app := &Server{
		server: grpc.NewServer(),
		doc: doc,
		client: c,
	}
	return app.init()
}

func (g *Server) init() *Server {
	logger.Info("initializing gRPC servers...")

	chainpb.RegisterChainServiceServer(g.server, NewChainServer(g.doc))
	accountpb.RegisterAccountServiceServer(g.server, NewAccountServer(g.client))
	blockpb.RegisterBlockServiceServer(g.server, NewBlockServer(g.client))
	statepb.RegisterStateServiceServer(g.server, NewStateServer(g.client))
	validatorpb.RegisterValidatorServiceServer(g.server, NewValidatorServer(g.client))
	transactionpb.RegisterTransactionServiceServer(g.server, NewTransactionServer(g.client))

	// Register reflection service on gRPC accountServer.
	reflection.Register(g.server)

	return g
}

func (g *Server) Start(serverAddress string, serverPort int64) error {
	logger.Info(fmt.Sprintf("starting grpc accountServer %s:%d...", serverAddress, serverPort), logger.Field("app", "grpc"))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverAddress, serverPort))
	if err != nil {
		return err
	}
	return g.server.Serve(lis)
}
