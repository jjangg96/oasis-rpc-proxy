package main

import (
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockserver"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/stateserver"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionserver"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorserver"
	"github.com/figment-networks/oasis-rpc-proxy/log"
	genesisFile "github.com/oasislabs/oasis-core/go/genesis/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	log.Info("initializing Oasis genesis document")
	if err := initGenesis(); err != nil {
		panic(err)
	}

	fmt.Println("State Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Error("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	blockpb.RegisterBlockServiceServer(s, blockserver.New())
	statepb.RegisterStateServiceServer(s, stateserver.New())
	validatorpb.RegisterValidatorServiceServer(s, validatorserver.New())
	transactionpb.RegisterTransactionServiceServer(s, transactionserver.New())

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Error("failed to serve: %v", err)
	}
}

func initGenesis() error {
	genesis, err := genesisFile.NewFileProvider("genesis.json")
	if err != nil {
		log.Error("failed to load genesis file", err)
		return err
	}

	// Retrieve the genesis document and use it to configure the ChainID for
	// signature domain separation. We do this as early as possible.
	doc, err := genesis.GetGenesisDocument()
	if err != nil {
		log.Error("failed to retrieve genesis document", err)
		return err
	}

	fmt.Printf("Chain context: '%v'\n\n", doc.ChainContext())
	fmt.Printf("Chain ID: '%v'\n\n", doc.Hash().String())
	var some []byte
	hash := doc.Hash()
	err2 := hash.UnmarshalBinary(some)
	if err2 != nil {

	}
	fmt.Printf("Chain ID: '%v'\n\n", some)

	doc.SetChainContext()

	return nil
}