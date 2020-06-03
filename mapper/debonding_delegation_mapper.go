package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/debondingdelegation/debondingdelegationpb"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	"github.com/oasislabs/oasis-core/go/staking/api"
)

func DebondingDelegationToPb(rawDebondingDelegations map[signature.PublicKey][]*api.DebondingDelegation) map[string]*debondingdelegationpb.DebondingDelegationInnerEntry {
	innerEntries := map[string]*debondingdelegationpb.DebondingDelegationInnerEntry{}
	for escrowPublicKey, innerItems := range rawDebondingDelegations {
		var dds []*debondingdelegationpb.DebondingDelegation
		for _, item := range innerItems {
			dds = append(dds, &debondingdelegationpb.DebondingDelegation{
				Shares:        item.Shares.ToBigInt().Bytes(),
				DebondEndTime: uint64(item.DebondEndTime),
			})
		}

		innerEntries[escrowPublicKey.String()] = &debondingdelegationpb.DebondingDelegationInnerEntry{
			DebondingDelegations: dds,
		}
	}
	return innerEntries
}


