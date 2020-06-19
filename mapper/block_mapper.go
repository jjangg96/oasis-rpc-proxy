package mapper

import (
	"encoding/base64"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/oasisprotocol/oasis-core/go/common/cbor"
	"github.com/oasisprotocol/oasis-core/go/consensus/api"
	tmApi "github.com/oasisprotocol/oasis-core/go/consensus/tendermint/api"
)

func BlockToPb(rawBlock api.Block) (*blockpb.Block, error) {
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
	for i, vote := range blockMeta.LastCommit.Signatures {
		// it has nil value when validator has not voted meaning that it was most likely offline
		votes = append(votes, &blockpb.Vote{
			BlockIdFlag:      int64(vote.BlockIDFlag),
			ValidatorAddress: vote.ValidatorAddress.String(),
			ValidatorIndex:   int64(i),
			Signature:        base64.StdEncoding.EncodeToString(vote.Signature),
		})
	}

	return &blockpb.Block{
		Header: &blockpb.Header{
			Version:            version,
			ChainId:            blockMeta.Header.ChainID,
			Height:             blockMeta.Header.Height,
			Time:               time,
			LastBlockId:        lastBlockId,
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
			BlockId: lastBlockId,
			Height:  blockMeta.LastCommit.Height,
			Round:   int64(blockMeta.LastCommit.Round),
			Hash:    blockMeta.LastCommit.Hash().String(),
			Votes:   votes,
		},
	}, nil
}
