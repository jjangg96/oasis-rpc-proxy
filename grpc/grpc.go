package main

import (
	"fmt"
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
	"github.com/figment-networks/oasis-rpc-proxy/log"
	"github.com/oasislabs/oasis-core/go/genesis/api"
	genesisFile "github.com/oasislabs/oasis-core/go/genesis/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	log.Info("initializing Oasis genesis document")
	doc := initGenesis()

	log.Info("starting grpc server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Error("failed to listen: %v", err)
		panic(err)
	}

	s := grpc.NewServer()
	chainpb.RegisterChainServiceServer(s, chainserver.New(*doc))
	blockpb.RegisterBlockServiceServer(s, blockserver.New())
	statepb.RegisterStateServiceServer(s, stateserver.New())
	validatorpb.RegisterValidatorServiceServer(s, validatorserver.New())
	transactionpb.RegisterTransactionServiceServer(s, transactionserver.New())

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Error("failed to serve: %v", err)
		panic(err)
	}
}

func initGenesis() *api.Document {
	genesis, err := genesisFile.NewFileProvider("genesis.json")
	if err != nil {
		log.Error("failed to load genesis file", err)
		panic(err)
	}

	// Retrieve the genesis document and use it to configure the ChainID for
	// signature domain separation. We do this as early as possible.
	doc, err := genesis.GetGenesisDocument()
	if err != nil {
		log.Error("failed to retrieve genesis document", err)
		panic(err)
	}

	log.Info(fmt.Sprintf("Chain context: '%v'", doc.ChainContext()))

	doc.SetChainContext()

	return doc
}