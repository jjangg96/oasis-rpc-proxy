package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/event/eventpb"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

func AddEscrowEventToPb(escrowEvent *api.AddEscrowEvent) *eventpb.AddEscrowEvent {
	return &eventpb.AddEscrowEvent{
		Owner:  escrowEvent.Owner.String(),
		Escrow: escrowEvent.Escrow.String(),
		Amount: escrowEvent.Amount.ToBigInt().Bytes(),
	}
}
