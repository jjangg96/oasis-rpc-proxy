package mapper

import (
	"github.com/jjangg96/oasis-rpc-proxy/grpc/debondingdelegation/debondingdelegationpb"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

func DebondingDelegationToPb(rawDebondingDelegations map[api.Address][]*api.DebondingDelegation) map[string]*debondingdelegationpb.DebondingDelegationInnerEntry {
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
