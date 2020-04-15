package schedulermapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/oasislabs/oasis-core/go/scheduler/api"
)

func ToPb(rawScheduler api.Genesis) *statepb.Scheduler {
	return &statepb.Scheduler{
		Params: &statepb.SchedulerParams{
			MinValidators:                int64(rawScheduler.Parameters.MinValidators),
			MaxValidators:                int64(rawScheduler.Parameters.MaxValidators),
			MaxValidatorsPerEntity:       int64(rawScheduler.Parameters.MaxValidatorsPerEntity),
			DebugBypassStake:             rawScheduler.Parameters.DebugBypassStake,
			DebugStaticValidators:        rawScheduler.Parameters.DebugStaticValidators,
			RewardFactorEpochElectionAny: rawScheduler.Parameters.RewardFactorEpochElectionAny.ToBigInt().Bytes(),
		},
	}
}
