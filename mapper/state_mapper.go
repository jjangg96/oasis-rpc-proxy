package mapper

import (
	"errors"

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

	// Compute consensus evidence parameters based on debonding period.
	debondingInterval := int64(rawState.Staking.Parameters.DebondingInterval)
	if debondingInterval == 0 {
		return nil, errors.New("debonding interval cannot equal zero")
	}
	epochInterval := rawState.EpochTime.Parameters.Interval
	if epochInterval == 0 {
		return nil, errors.New("epoch interval cannot equal zero")
	}

	maxAgeNumBlocks := debondingInterval * epochInterval

	state := &statepb.State{
		Height:    rawState.Height,
		Time:      time,
		ChainID:   rawState.ChainID,
		Registry:  RegistryToPb(rawState.Registry),
		Staking:   StakingToPb(rawState.Staking),
		Scheduler: SchedulerToPb(rawState.Scheduler),
		Consensus: ConsensusToPb(rawState.Consensus, maxAgeNumBlocks),
	}
	return state, nil
}
