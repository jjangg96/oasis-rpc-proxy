package mapper

import (
	"github.com/jjangg96/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

func AccountToPb(rawAccount api.Account) *accountpb.Account {
	// Rates
	var rates []*accountpb.CommissionRateStep
	for _, rate := range rawAccount.Escrow.CommissionSchedule.Rates {
		rates = append(rates, &accountpb.CommissionRateStep{
			Start: uint64(rate.Start),
			Rate:  rate.Rate.ToBigInt().Bytes(),
		})
	}

	// Bounds
	var bounds []*accountpb.CommissionRateBoundStep
	for _, bound := range rawAccount.Escrow.CommissionSchedule.Bounds {
		bounds = append(bounds, &accountpb.CommissionRateBoundStep{
			Start:   uint64(bound.Start),
			RateMin: bound.RateMin.ToBigInt().Bytes(),
			RateMax: bound.RateMax.ToBigInt().Bytes(),
		})
	}

	// Claims
	claims := map[string]*accountpb.ThresholdKinds{}
	for claim, rawThresholds := range rawAccount.Escrow.StakeAccumulator.Claims {
		var kinds []*accountpb.StakeThreshold
		for _, threshold := range rawThresholds {
			stakeThreshold := &accountpb.StakeThreshold{}

			if threshold.Global != nil {
				stakeThreshold.Global = threshold.Global.String()
			}

			if threshold.Constant != nil {
				stakeThreshold.Constant = threshold.Constant.ToBigInt().Bytes()
			}

			kinds = append(kinds, stakeThreshold)
		}

		claims[string(claim)] = &accountpb.ThresholdKinds{
			Kinds: kinds,
		}
	}

	return &accountpb.Account{
		General: &accountpb.GeneralAccount{
			Balance: rawAccount.General.Balance.ToBigInt().Bytes(),
			Nonce:   rawAccount.General.Nonce,
		},
		Escrow: &accountpb.EscrowAccount{
			Active: &accountpb.SharePool{
				Balance:     rawAccount.Escrow.Active.Balance.ToBigInt().Bytes(),
				TotalShares: rawAccount.Escrow.Active.TotalShares.ToBigInt().Bytes(),
			},
			Debonding: &accountpb.SharePool{
				Balance:     rawAccount.Escrow.Debonding.Balance.ToBigInt().Bytes(),
				TotalShares: rawAccount.Escrow.Debonding.TotalShares.ToBigInt().Bytes(),
			},
			CommissionSchedule: &accountpb.CommissionSchedule{
				Rates:  rates,
				Bounds: bounds,
			},
			StakeAccumulator: &accountpb.StakeAccumulator{
				Claims: claims,
			},
		},
	}
}
