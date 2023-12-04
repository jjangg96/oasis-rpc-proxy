package mapper

import (
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/oasisprotocol/oasis-core/go/consensus/genesis"
)

func ConsensusToPb(rawConsensus genesis.Genesis, maxAgeNumBlocks int64) *statepb.Consensus {
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
			MaxTxSize:    rawConsensus.Parameters.MaxTxSize,
			MaxBlockSize: rawConsensus.Parameters.MaxBlockSize,
			MaxBlockGas:  uint64(rawConsensus.Parameters.MaxBlockGas),

			MaxEvidenceAgeBlocks: uint64(maxAgeNumBlocks),
			MaxEvidenceAgeTime:   (time.Duration(maxAgeNumBlocks) * (rawConsensus.Parameters.TimeoutCommit + 1*time.Second)).String(),
		},
	}
}
