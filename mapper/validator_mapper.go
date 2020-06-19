package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/oasisprotocol/oasis-core/go/common/node"
	tmcrypto "github.com/oasisprotocol/oasis-core/go/consensus/tendermint/crypto"
	"github.com/oasisprotocol/oasis-core/go/scheduler/api"
	stakingApi "github.com/oasisprotocol/oasis-core/go/staking/api"
)

func ValidatorToPb(validator *api.Validator, node *node.Node) *validatorpb.Validator {
	cID := node.Consensus.ID
	tmAddr := tmcrypto.PublicKeyToTendermint(&cID).Address().String()

	// P2P addresses
	var p2pAddresses []string
	for _, pa := range node.P2P.Addresses {
		p2pAddresses = append(p2pAddresses, pa.String())
	}

	// Consensus addresses
	var consensusAddresses []*validatorpb.ConsensusAddress
	for _, ca := range node.Consensus.Addresses {
		consensusAddresses = append(consensusAddresses, &validatorpb.ConsensusAddress{
			Id:      ca.ID.String(),
			Address: ca.Address.String(),
		})
	}

	// Runtimes
	var runtimes []*validatorpb.Runtime
	for _, r := range node.Runtimes {
		runtimes = append(runtimes, &validatorpb.Runtime{
			Id: r.ID.String(),
			Version: &validatorpb.Version{
				Major: uint32(r.Version.Major),
				Minor: uint32(r.Version.Minor),
				Patch: uint32(r.Version.Patch),
			},
		})
	}

	return &validatorpb.Validator{
		Id:                validator.ID.String(),
		Address:           stakingApi.NewAddress(node.EntityID).String(),
		TendermintAddress: tmAddr,
		VotingPower:       validator.VotingPower,

		Node: &validatorpb.Node{
			Id:         node.ID.String(),
			EntityId:   node.EntityID.String(),
			Expiration: node.Expiration,

			P2P: &validatorpb.P2PInfo{
				Id:        node.P2P.ID.String(),
				Addresses: p2pAddresses,
			},

			Consensus: &validatorpb.ConsensusInfo{
				Id:        node.Consensus.ID.String(),
				Addresses: consensusAddresses,
			},

			Runtimes: runtimes,
			Roles:    uint32(node.Roles),
		},
	}
}
