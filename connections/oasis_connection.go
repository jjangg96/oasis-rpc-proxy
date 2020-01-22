package connections

import (
	"github.com/figment-networks/oasis-rpc-proxy/config"
	oasisGrpc "github.com/oasislabs/oasis-core/go/common/grpc"
	"google.golang.org/grpc"
)

func GetOasisConn() (*grpc.ClientConn, error) {
	return oasisGrpc.Dial(
		config.GetOasisSocket(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	)
}
