package consensusmapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/oasislabs/oasis-core/go/consensus/genesis"
)

func ToPb(rawConsensus genesis.Genesis) *statepb.Consensus {
	return &statepb.Consensus{
		Backend: rawConsensus.Backend,
		Params: &statepb.ConsensusParams{
			TimeoutCommit: &duration.Duration{
				Seconds: int64(rawConsensus.Parameters.TimeoutCommit.Seconds()),
				Nanos:   int32(rawConsensus.Parameters.TimeoutCommit.Nanoseconds()),
			},
			SkipTimeoutCommit: rawConsensus.Parameters.SkipTimeoutCommit,
			EmptyBlockInterval: &duration.Duration{
				Seconds: int64(rawConsensus.Parameters.EmptyBlockInterval.Seconds()),
				Nanos:   int32(rawConsensus.Parameters.EmptyBlockInterval.Nanoseconds()),
			},
			MaxTxSize:            rawConsensus.Parameters.MaxTxSize,
			MaxBlockSize:         rawConsensus.Parameters.MaxBlockSize,
			MaxBlockGas:          uint64(rawConsensus.Parameters.MaxBlockGas),
			MaxEvidenceAgeBlocks: rawConsensus.Parameters.MaxEvidenceAgeBlocks,
			MaxEvidenceAgeTime:   rawConsensus.Parameters.MaxEvidenceAgeTime.String(),
		},
	}
}
