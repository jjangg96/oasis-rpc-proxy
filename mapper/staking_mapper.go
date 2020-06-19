package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/debondingdelegation/debondingdelegationpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/delegation/delegationpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

func StakingToPb(rawStaking api.Genesis) *statepb.Staking {
	// Thresholds
	thresholds := map[int64][]byte{}
	for kind, quantity := range rawStaking.Parameters.Thresholds {
		thresholds[int64(kind)] = quantity.ToBigInt().Bytes()
	}

	// Reward Schedule
	var rewardSchedule []*statepb.RewardStep
	for _, step := range rawStaking.Parameters.RewardSchedule {
		rewardSchedule = append(rewardSchedule, &statepb.RewardStep{
			Scale: step.Scale.ToBigInt().Bytes(),
			Until: uint64(step.Until),
		})
	}

	// Slashing
	slashing := map[string]*statepb.Slash{}
	for reason, slash := range rawStaking.Parameters.Slashing {
		slashing[reason.String()] = &statepb.Slash{
			Amount:         slash.Amount.ToBigInt().Bytes(),
			FreezeInterval: uint64(slash.FreezeInterval),
		}
	}

	// Gas costs
	gasCosts := map[string]uint64{}
	for op, gas := range rawStaking.Parameters.GasCosts {
		gasCosts[string(op)] = uint64(gas)
	}

	// Undisable transfers from
	undisableTransfersFrom := map[string]bool{}
	for key, b := range rawStaking.Parameters.UndisableTransfersFrom {
		undisableTransfersFrom[key.String()] = b
	}

	// Ledger
	ledger := map[string]*accountpb.Account{}
	for key, account := range rawStaking.Ledger {
		ledger[key.String()] = AccountToPb(*account)
	}

	// Delegations
	delegations := map[string]*delegationpb.DelegationEntry{}
	for validatorId, items := range rawStaking.Delegations {
		delegations[validatorId.String()] = &delegationpb.DelegationEntry{
			Entries: DelegationToPb(items),
		}
	}

	// Debonding delegations
	debondingDelegations := map[string]*debondingdelegationpb.DebondingDelegationEntry{}
	for validatorId, items := range rawStaking.DebondingDelegations {
		debondingDelegations[validatorId.String()] = &debondingdelegationpb.DebondingDelegationEntry{
			Entries: DebondingDelegationToPb(items),
		}
	}

	return &statepb.Staking{
		TotalSupply: rawStaking.TotalSupply.ToBigInt().Bytes(),
		CommonPool:  rawStaking.CommonPool.ToBigInt().Bytes(),
		Parameters: &statepb.StakingParameters{
			Thresholds:                        thresholds,
			DebondingInterval:                 uint64(rawStaking.Parameters.DebondingInterval),
			RewardSchedule:                    rewardSchedule,
			SigningRewardThresholdNumerator:   rawStaking.Parameters.SigningRewardThresholdNumerator,
			SigningRewardThresholdDenominator: rawStaking.Parameters.SigningRewardThresholdDenominator,
			CommissionScheduleRules: &statepb.CommissionScheduleRules{
				RateBoundLead:      uint64(rawStaking.Parameters.CommissionScheduleRules.RateBoundLead),
				RateChangeInterval: uint64(rawStaking.Parameters.CommissionScheduleRules.RateChangeInterval),
				MaxBoundSteps:      int64(rawStaking.Parameters.CommissionScheduleRules.MaxBoundSteps),
				MaxRateSteps:       int64(rawStaking.Parameters.CommissionScheduleRules.MaxRateSteps),
			},
			Slashing:                  slashing,
			GasCosts:                  gasCosts,
			MinDelegationAmount:       rawStaking.Parameters.MinDelegationAmount.ToBigInt().Bytes(),
			DisableTransfers:          rawStaking.Parameters.DisableTransfers,
			DisableDelegation:         rawStaking.Parameters.DisableDelegation,
			UndisableTransfersFrom:    undisableTransfersFrom,
			FeeSplitWeightVote:        rawStaking.Parameters.FeeSplitWeightVote.ToBigInt().Bytes(),
			FeeSplitWeightPropose:     rawStaking.Parameters.FeeSplitWeightPropose.ToBigInt().Bytes(),
			FeeSplitWeightNextPropose: rawStaking.Parameters.FeeSplitWeightNextPropose.ToBigInt().Bytes(),
			RewardFactorEpochSigned:   rawStaking.Parameters.RewardFactorEpochSigned.ToBigInt().Bytes(),
			RewardFactorBlockProposed: rawStaking.Parameters.RewardFactorBlockProposed.ToBigInt().Bytes(),
		},
		Ledger:               ledger,
		Delegations:          delegations,
		DebondingDelegations: debondingDelegations,
	}
}
