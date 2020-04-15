package statemapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/figment-networks/oasis-rpc-proxy/mappers/consensusmapper"
	"github.com/figment-networks/oasis-rpc-proxy/mappers/registrymapper"
	"github.com/figment-networks/oasis-rpc-proxy/mappers/schedulermapper"
	"github.com/figment-networks/oasis-rpc-proxy/mappers/stakingmapper"
	"github.com/golang/protobuf/ptypes"
	genesisApi "github.com/oasislabs/oasis-core/go/genesis/api"
)

func ToPb(rawState *genesisApi.Document) (*statepb.State, error) {
	// Time
	time, err := ptypes.TimestampProto(rawState.Time)
	if err != nil {
		return nil, err
	}

	state := &statepb.State{
		Height:    rawState.Height,
		Time:      time,
		ChainID:   rawState.ChainID,
		Registry:  registrymapper.ToPb(rawState.Registry),
		Staking:   stakingmapper.ToPb(rawState.Staking),
		Scheduler: schedulermapper.ToPb(rawState.Scheduler),
		Consensus: consensusmapper.ToPb(rawState.Consensus),
	}
	return state, nil
}
