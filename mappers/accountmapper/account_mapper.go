package accountmapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/oasislabs/oasis-core/go/staking/api"
)

func ToPb(rawAccount api.Account) *statepb.Account {
	// Rates
	var rates []*statepb.CommissionRateStep
	for _, rate := range rawAccount.Escrow.CommissionSchedule.Rates {
		rates = append(rates, &statepb.CommissionRateStep{
			Start: uint64(rate.Start),
			Rate:  rate.Rate.ToBigInt().Bytes(),
		})
	}

	// Bounds
	var bounds []*statepb.CommissionRateBoundStep
	for _, bound := range rawAccount.Escrow.CommissionSchedule.Bounds {
		bounds = append(bounds, &statepb.CommissionRateBoundStep{
			Start:   uint64(bound.Start),
			RateMin: bound.RateMin.ToBigInt().Bytes(),
			RateMax: bound.RateMax.ToBigInt().Bytes(),
		})
	}

	// Claims
	claims := map[string]*statepb.ThresholdKinds{}
	for claim, rawKinds := range rawAccount.Escrow.StakeAccumulator.Claims {
		var kinds []int32
		for _, kind := range rawKinds {
			kinds = append(kinds, int32(kind))
		}

		claims[string(claim)] = &statepb.ThresholdKinds{
			Kinds: kinds,
		}
	}

	return &statepb.Account{
		General: &statepb.GeneralAccount{
			Balance:            rawAccount.General.Balance.ToBigInt().Bytes(),
			Nonce:              rawAccount.General.Nonce,
		},
		Escrow: &statepb.EscrowAccount{
			Active: &statepb.SharePool{
				Balance:     rawAccount.Escrow.Active.Balance.ToBigInt().Bytes(),
				TotalShares: rawAccount.Escrow.Active.TotalShares.ToBigInt().Bytes(),
			},
			Debonding: &statepb.SharePool{
				Balance:     rawAccount.Escrow.Debonding.Balance.ToBigInt().Bytes(),
				TotalShares: rawAccount.Escrow.Debonding.TotalShares.ToBigInt().Bytes(),
			},
			CommissionSchedule: &statepb.CommissionSchedule{
				Rates:  rates,
				Bounds: bounds,
			},
			StakeAccumulator: &statepb.StakeAccumulator{
				Claims: claims,
			},
		},
	}
}
