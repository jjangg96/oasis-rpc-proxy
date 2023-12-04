package mapper

import (
	"github.com/jjangg96/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/oasisprotocol/oasis-core/go/scheduler/api"
)

func SchedulerToPb(rawScheduler api.Genesis) *statepb.Scheduler {
	return &statepb.Scheduler{
		Params: &statepb.SchedulerParams{
			MinValidators:                int64(rawScheduler.Parameters.MinValidators),
			MaxValidators:                int64(rawScheduler.Parameters.MaxValidators),
			MaxValidatorsPerEntity:       int64(rawScheduler.Parameters.MaxValidatorsPerEntity),
			DebugBypassStake:             rawScheduler.Parameters.DebugBypassStake,
			RewardFactorEpochElectionAny: rawScheduler.Parameters.RewardFactorEpochElectionAny.ToBigInt().Bytes(),
		},
	}
}
