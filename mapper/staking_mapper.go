package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/debondingdelegation/debondingdelegationpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/delegation/delegationpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

func StakingToPb(rawStaking api.Genesis) *statepb.Staking {
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
		TotalSupply:          rawStaking.TotalSupply.ToBigInt().Bytes(),
		CommonPool:           rawStaking.CommonPool.ToBigInt().Bytes(),
		Parameters:           ConsensusParametersToStakingParameters(rawStaking.Parameters),
		Ledger:               ledger,
		Delegations:          delegations,
		DebondingDelegations: debondingDelegations,
	}
}

func ConsensusParametersToStakingParameters(parameters api.ConsensusParameters) *statepb.StakingParameters {
	// Thresholds
	thresholds := map[int64][]byte{}
	for kind, quantity := range parameters.Thresholds {
		thresholds[int64(kind)] = quantity.ToBigInt().Bytes()
	}

	// Reward Schedule
	var rewardSchedule []*statepb.RewardStep
	for _, step := range parameters.RewardSchedule {
		rewardSchedule = append(rewardSchedule, &statepb.RewardStep{
			Scale: step.Scale.ToBigInt().Bytes(),
			Until: uint64(step.Until),
		})
	}

	// Slashing
	slashing := map[string]*statepb.Slash{}
	for reason, slash := range parameters.Slashing {
		slashing[reason.String()] = &statepb.Slash{
			Amount:         slash.Amount.ToBigInt().Bytes(),
			FreezeInterval: uint64(slash.FreezeInterval),
		}
	}

	// Gas costs
	gasCosts := map[string]uint64{}
	for op, gas := range parameters.GasCosts {
		gasCosts[string(op)] = uint64(gas)
	}

	// Undisable transfers from
	undisableTransfersFrom := map[string]bool{}
	for key, b := range parameters.UndisableTransfersFrom {
		undisableTransfersFrom[key.String()] = b
	}

	return &statepb.StakingParameters{
		Thresholds:                        thresholds,
		DebondingInterval:                 uint64(parameters.DebondingInterval),
		RewardSchedule:                    rewardSchedule,
		SigningRewardThresholdNumerator:   parameters.SigningRewardThresholdNumerator,
		SigningRewardThresholdDenominator: parameters.SigningRewardThresholdDenominator,
		CommissionScheduleRules: &statepb.CommissionScheduleRules{
			RateBoundLead:      uint64(parameters.CommissionScheduleRules.RateBoundLead),
			RateChangeInterval: uint64(parameters.CommissionScheduleRules.RateChangeInterval),
			MaxBoundSteps:      int64(parameters.CommissionScheduleRules.MaxBoundSteps),
			MaxRateSteps:       int64(parameters.CommissionScheduleRules.MaxRateSteps),
		},
		Slashing:                  slashing,
		GasCosts:                  gasCosts,
		MinDelegationAmount:       parameters.MinDelegationAmount.ToBigInt().Bytes(),
		DisableTransfers:          parameters.DisableTransfers,
		DisableDelegation:         parameters.DisableDelegation,
		UndisableTransfersFrom:    undisableTransfersFrom,
		FeeSplitWeightVote:        parameters.FeeSplitWeightVote.ToBigInt().Bytes(),
		FeeSplitWeightPropose:     parameters.FeeSplitWeightPropose.ToBigInt().Bytes(),
		FeeSplitWeightNextPropose: parameters.FeeSplitWeightNextPropose.ToBigInt().Bytes(),
		RewardFactorEpochSigned:   parameters.RewardFactorEpochSigned.ToBigInt().Bytes(),
		RewardFactorBlockProposed: parameters.RewardFactorBlockProposed.ToBigInt().Bytes(),
	}
}
