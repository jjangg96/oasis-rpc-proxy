package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/event/eventpb"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

func EventToPb(rawEvent *api.Event) *eventpb.AddEscrowEvent {
	return &eventpb.AddEscrowEvent{
		Owner:  rawEvent.Escrow.Add.Owner.String(),
		Escrow: rawEvent.Escrow.Add.Escrow.String(),
		Amount: rawEvent.Escrow.Add.Amount.ToBigInt().Bytes(),
	}
}
