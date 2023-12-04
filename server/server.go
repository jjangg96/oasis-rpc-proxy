package server

import (
	"fmt"
	"net"

	"github.com/jjangg96/oasis-rpc-proxy/client"
	"github.com/jjangg96/oasis-rpc-proxy/config"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/chain/chainpb"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/debondingdelegation/debondingdelegationpb"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/delegation/delegationpb"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/event/eventpb"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/jjangg96/oasis-rpc-proxy/metric"
	"github.com/jjangg96/oasis-rpc-proxy/utils/logger"
	"github.com/oasisprotocol/oasis-core/go/genesis/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server *grpc.Server

	cfg    *config.Config
	client *client.Client
	doc    *api.Document
}

func New(cfg *config.Config, c *client.Client, doc *api.Document) *Server {

	app := &Server{
		server: grpc.NewServer(
			grpc.MaxSendMsgSize(cfg.GrpcMaxSendSize),
			grpc.MaxRecvMsgSize(cfg.GrpcMaxRecvSize),
		),
		cfg:    cfg,
		client: c,
		doc:    doc,
	}
	return app.init()
}

func (s *Server) init() *Server {
	logger.Info("initializing gRPC servers...", logger.Field("app", "server"))

	chainpb.RegisterChainServiceServer(s.server, NewChainServer(s.client, s.doc))
	accountpb.RegisterAccountServiceServer(s.server, NewAccountServer(s.client))
	blockpb.RegisterBlockServiceServer(s.server, NewBlockServer(s.client))
	eventpb.RegisterEventServiceServer(s.server, NewEventServer(s.client))
	statepb.RegisterStateServiceServer(s.server, NewStateServer(s.client))
	validatorpb.RegisterValidatorServiceServer(s.server, NewValidatorServer(s.client))
	transactionpb.RegisterTransactionServiceServer(s.server, NewTransactionServer(s.client))
	delegationpb.RegisterDelegationServiceServer(s.server, NewDelegationServer(s.client))
	debondingdelegationpb.RegisterDebondingDelegationServiceServer(s.server, NewDebondingDelegationServer(s.client))

	// Register reflection service on gRPC accountServer.
	reflection.Register(s.server)

	return s
}

func (s *Server) Start(serverAddress string, serverPort int64) error {
	logger.Info(fmt.Sprintf("starting grpc accountServer %s:%d...", serverAddress, serverPort), logger.Field("app", "server"))

	go s.startMetricsServer()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverAddress, serverPort))
	if err != nil {
		return err
	}
	return s.server.Serve(lis)
}

func (s *Server) startMetricsServer() error {
	return metric.NewClientMetric().StartServer(s.cfg.ServerMetricAddr, s.cfg.MetricServerUrl)
}
