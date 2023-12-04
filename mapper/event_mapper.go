package mapper

import (
	"github.com/jjangg96/oasis-rpc-proxy/grpc/event/eventpb"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

func AddEscrowEventToPb(escrowEvent *api.AddEscrowEvent) *eventpb.AddEscrowEvent {
	return &eventpb.AddEscrowEvent{
		Owner:     escrowEvent.Owner.String(),
		Escrow:    escrowEvent.Escrow.String(),
		Amount:    escrowEvent.Amount.ToBigInt().Bytes(),
		NewShares: escrowEvent.NewShares.ToBigInt().Bytes(),
	}
}

func TakeEscrowEventToPb(escrowEvent *api.TakeEscrowEvent) *eventpb.TakeEscrowEvent {
	return &eventpb.TakeEscrowEvent{
		Owner:  escrowEvent.Owner.String(),
		Amount: escrowEvent.Amount.ToBigInt().Bytes(),
	}
}

func TransferEventToPb(ev *api.TransferEvent) *eventpb.TransferEvent {
	return &eventpb.TransferEvent{
		From:   ev.From.String(),
		To:     ev.To.String(),
		Amount: ev.Amount.ToBigInt().Bytes(),
	}
}
