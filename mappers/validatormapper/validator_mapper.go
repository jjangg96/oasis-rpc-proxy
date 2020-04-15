package validatormapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/oasislabs/oasis-core/go/common/node"
	tmcrypto "github.com/oasislabs/oasis-core/go/consensus/tendermint/crypto"
	"github.com/oasislabs/oasis-core/go/scheduler/api"
)

func ToPb(validator *api.Validator, node *node.Node) *validatorpb.Validator {
	cID := node.Consensus.ID
	tmAddr := tmcrypto.PublicKeyToTendermint(&cID).Address().String()

	// Committee addresses
	var committeeAddresses []*validatorpb.CommitteeAddress
	for _, ca := range node.Committee.Addresses {
		committeeAddresses = append(committeeAddresses, &validatorpb.CommitteeAddress{
			Certificate: ca.Certificate,
			Address:     ca.Address.String(),
		})
	}

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
		Id:          validator.ID.String(),
		Address:     tmAddr,
		VotingPower: validator.VotingPower,

		Node: &validatorpb.Node{
			Id:         node.ID.String(),
			EntityId:   node.EntityID.String(),
			Expiration: node.Expiration,

			Committee: &validatorpb.CommitteeInfo{
				Certificate: node.Committee.Certificate,
				Addresses:   committeeAddresses,
			},

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
