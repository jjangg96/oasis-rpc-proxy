package chainmapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/chain/chainpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/oasislabs/oasis-core/go/genesis/api"
)

func ToPb(doc api.Document) (*chainpb.Chain, error) {
	time, err := ptypes.TimestampProto(doc.Time)
	if err != nil {
		return nil, err
	}

	return &chainpb.Chain{
		Id:     doc.Hash().String(),
		Name:   doc.ChainID,
		GenesisTime: time,
		Height:      doc.Height,
	}, nil
}
