package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/golang/protobuf/ptypes"
	genesisApi "github.com/oasisprotocol/oasis-core/go/genesis/api"
)

func StateToPb(rawState *genesisApi.Document) (*statepb.State, error) {
	// Time
	time, err := ptypes.TimestampProto(rawState.Time)
	if err != nil {
		return nil, err
	}

	state := &statepb.State{
		Height:    rawState.Height,
		Time:      time,
		ChainID:   rawState.ChainID,
		Registry:  RegistryToPb(rawState.Registry),
		Staking:   StakingToPb(rawState.Staking),
		Scheduler: SchedulerToPb(rawState.Scheduler),
		Consensus: ConsensusToPb(rawState.Consensus),
	}
	return state, nil
}
