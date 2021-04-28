package mapper

import (
	"errors"
	"fmt"

	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	beacon "github.com/oasisprotocol/oasis-core/go/beacon/api"
	cmdFlags "github.com/oasisprotocol/oasis-core/go/oasis-node/cmd/common/flags"
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
	// epochInterval := rawState.EpochTime.Parameters.Interval
	var epochInterval int64
	switch rawState.Beacon.Parameters.Backend {
	case beacon.BackendInsecure:
		params := rawState.Beacon.Parameters.InsecureParameters
		epochInterval = params.Interval
		if epochInterval == 0 && cmdFlags.DebugDontBlameOasis() && rawState.Beacon.Parameters.DebugMockBackend {
			// Use a default of 100 blocks in case epoch interval is unset
			// and we are using debug mode.
			epochInterval = 100
		}
	case beacon.BackendPVSS:
		// Note: This assumes no protocol failures (the common case).
		// In the event of a failure, it is entirely possible that epochs
		// can drag on for significantly longer.
		params := rawState.Beacon.Parameters.PVSSParameters
		epochInterval = params.CommitInterval + params.RevealInterval + params.TransitionDelay
	default:
		return nil, fmt.Errorf("tendermint: unknown beacon backend: '%s'", rawState.Beacon.Parameters.Backend)
	}
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
