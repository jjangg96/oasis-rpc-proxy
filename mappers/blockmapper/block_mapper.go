package blockmapper

import (
	"encoding/base64"
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/oasislabs/oasis-core/go/common/cbor"
	"github.com/oasislabs/oasis-core/go/consensus/api"
	tmApi "github.com/oasislabs/oasis-core/go/consensus/tendermint/api"
)

func ToPb(rawBlock api.Block) (*blockpb.Block, error) {
	var blockMeta tmApi.BlockMeta
	if err := cbor.Unmarshal(rawBlock.Meta, &blockMeta); err != nil {
		return nil, err
	}

	// Version
	version := &blockpb.Version{
		Block: blockMeta.Header.Version.Block.Uint64(),
		App:   blockMeta.Header.Version.App.Uint64(),
	}

	// LastBlockId
	lastBlockId := &blockpb.BlockID{
		Hash: blockMeta.Header.LastBlockID.Hash.String(),
	}

	// Time
	time, err := ptypes.TimestampProto(blockMeta.Header.Time)
	if err != nil {
		return nil, err
	}

	// Votes
	var votes []*blockpb.Vote
	for _, vote := range blockMeta.LastCommit.Precommits {
		// it has nil value when validator has not voted meaning that it was most likely offline
		if vote == nil {
			votes = append(votes, &blockpb.Vote{})
		} else {
			votes = append(votes, &blockpb.Vote{
				Type:                 fmt.Sprintf("%d", vote.Type),
				ValidatorAddress:     vote.ValidatorAddress.String(),
				ValidatorIndex:       int64(vote.ValidatorIndex),
				Signature:            base64.StdEncoding.EncodeToString(vote.Signature),
			})
		}
	}

	return &blockpb.Block{
		Header: &blockpb.Header{
			Version: version,
			ChainId:  blockMeta.Header.ChainID,
			Height:   blockMeta.Header.Height,
			Time:     time,
			NumTxs:   blockMeta.Header.NumTxs,
			TotalTxs: blockMeta.Header.TotalTxs,
			LastBlockId: lastBlockId,
			LastCommitHash:     blockMeta.Header.LastCommitHash.String(),
			DataHash:           blockMeta.Header.DataHash.String(),
			ValidatorsHash:     blockMeta.Header.ValidatorsHash.String(),
			NextValidatorsHash: blockMeta.Header.NextValidatorsHash.String(),
			ConsensusHash:      blockMeta.Header.ConsensusHash.String(),
			AppHash:            blockMeta.Header.AppHash.String(),
			LastResultsHash:    blockMeta.Header.LastResultsHash.String(),
			EvidenceHash:       blockMeta.Header.EvidenceHash.String(),
			ProposerAddress:    blockMeta.Header.ProposerAddress.String(),
		},
		LastCommit: &blockpb.Commit{
			BlockId:              lastBlockId,
			Height:               blockMeta.LastCommit.Height(),
			Round:                int64(blockMeta.LastCommit.Round()),
			Hash:                 blockMeta.LastCommit.Hash().String(),
			Votes:                votes,
		},
	}, nil
}
