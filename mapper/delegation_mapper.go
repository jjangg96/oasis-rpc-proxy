package mapper

import (
	"github.com/jjangg96/oasis-rpc-proxy/grpc/delegation/delegationpb"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

func DelegationToPb(rawDelegations map[api.Address]*api.Delegation) map[string]*delegationpb.Delegation {
	entries := map[string]*delegationpb.Delegation{}
	for escrowPublicKey, delegation := range rawDelegations {
		entries[escrowPublicKey.String()] = &delegationpb.Delegation{
			Shares: delegation.Shares.ToBigInt().Bytes(),
		}
	}

	return entries
}
