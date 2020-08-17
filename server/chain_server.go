package server

import (
	"context"

	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/chain/chainpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
	"github.com/golang/protobuf/ptypes"
	"github.com/oasisprotocol/oasis-core/go/genesis/api"
	stakingApi "github.com/oasisprotocol/oasis-core/go/staking/api"
	tmtypes "github.com/tendermint/tendermint/types"
)

type ChainServer interface {
	GetStatus(context.Context, *chainpb.GetStatusRequest) (*chainpb.GetStatusResponse, error)
	GetMetaByHeight(context.Context, *chainpb.GetMetaByHeightRequest) (*chainpb.GetMetaByHeightResponse, error)
	GetHead(context.Context, *chainpb.GetHeadRequest) (*chainpb.GetHeadResponse, error)
	GetConstants(context.Context, *chainpb.GetConstantsRequest) (*chainpb.GetConstantsResponse, error)
}

type chainServer struct {
	client *client.Client
	doc    *api.Document
}

func NewChainServer(c *client.Client, doc *api.Document) ChainServer {
	return &chainServer{
		client: c,
		doc:    doc,
	}
}

func (s *chainServer) GetConstants(ctx context.Context, req *chainpb.GetConstantsRequest) (*chainpb.GetConstantsResponse, error) {
	return &chainpb.GetConstantsResponse{CommonPoolAddress: stakingApi.CommonPoolAddress.String()}, nil
}

func (s *chainServer) GetStatus(ctx context.Context, req *chainpb.GetStatusRequest) (*chainpb.GetStatusResponse, error) {
	time, err := ptypes.TimestampProto(s.doc.Time)
	if err != nil {
		return nil, err
	}

	return &chainpb.GetStatusResponse{
		Id:            s.doc.Hash().String()[:tmtypes.MaxChainIDLen],
		Name:          s.doc.ChainID,
		GenesisTime:   time,
		GenesisHeight: s.doc.Height,
	}, nil
}

func (s *chainServer) GetMetaByHeight(ctx context.Context, req *chainpb.GetMetaByHeightRequest) (*chainpb.GetMetaByHeightResponse, error) {
	block, err := s.getBlock(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	epochTime, err := s.client.Consensus.GetEpochByHeight(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	return &chainpb.GetMetaByHeightResponse{
		Height:       block.GetHeader().GetHeight(),
		Time:         block.GetHeader().GetTime(),
		AppVersion:   block.GetHeader().GetVersion().GetApp(),
		BlockVersion: block.GetHeader().GetVersion().GetBlock(),
		Epoch:        uint64(epochTime),
	}, nil
}

func (s *chainServer) GetHead(ctx context.Context, req *chainpb.GetHeadRequest) (*chainpb.GetHeadResponse, error) {
	block, err := s.getBlock(ctx, 0)
	if err != nil {
		return nil, err
	}

	return &chainpb.GetHeadResponse{
		Height: block.GetHeader().GetHeight(),
		Time:   block.GetHeader().GetTime(),
	}, nil
}

func (s *chainServer) getBlock(ctx context.Context, height int64) (*blockpb.Block, error) {
	rawBlock, err := s.client.Consensus.GetBlockByHeight(ctx, height)
	if err != nil {
		return nil, err
	}

	block, err := mapper.BlockToPb(*rawBlock)
	if err != nil {
		return nil, err
	}
	return block, nil
}
