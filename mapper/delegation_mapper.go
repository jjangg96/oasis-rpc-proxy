package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/delegation/delegationpb"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	"github.com/oasislabs/oasis-core/go/staking/api"
)

func DelegationToPb(rawDelegations map[signature.PublicKey]*api.Delegation) map[string]*delegationpb.Delegation {
	entries := map[string]*delegationpb.Delegation{}
	for escrowPublicKey, delegation := range rawDelegations {
		entries[escrowPublicKey.String()] = &delegationpb.Delegation{
			Shares: delegation.Shares.ToBigInt().Bytes(),
		}
	}

	return entries
}

