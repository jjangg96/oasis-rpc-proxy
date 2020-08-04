package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/oasisprotocol/oasis-core/go/common/node"
	tmcrypto "github.com/oasisprotocol/oasis-core/go/consensus/tendermint/crypto"
	epochTimeApi "github.com/oasisprotocol/oasis-core/go/epochtime/api"
	"github.com/oasisprotocol/oasis-core/go/scheduler/api"
	stakingApi "github.com/oasisprotocol/oasis-core/go/staking/api"
	"math/big"
)

func ValidatorToPb(rawValidator *api.Validator, address string, rawNode *node.Node, rawAccount *stakingApi.Account, rawEpochTime epochTimeApi.EpochTime) *validatorpb.Validator {
	cID := rawNode.Consensus.ID
	tmAddr := tmcrypto.PublicKeyToTendermint(&cID).Address().String()
	rateNumerator := rawAccount.Escrow.CommissionSchedule.CurrentRate(rawEpochTime)

	// P2P addresses
	var p2pAddresses []string
	for _, pa := range rawNode.P2P.Addresses {
		p2pAddresses = append(p2pAddresses, pa.String())
	}

	// Consensus addresses
	var consensusAddresses []*validatorpb.ConsensusAddress
	for _, ca := range rawNode.Consensus.Addresses {
		consensusAddresses = append(consensusAddresses, &validatorpb.ConsensusAddress{
			Id:      ca.ID.String(),
			Address: ca.Address.String(),
		})
	}

	// Runtimes
	var runtimes []*validatorpb.Runtime
	for _, r := range rawNode.Runtimes {
		runtimes = append(runtimes, &validatorpb.Runtime{
			Id: r.ID.String(),
			Version: &validatorpb.Version{
				Major: uint32(r.Version.Major),
				Minor: uint32(r.Version.Minor),
				Patch: uint32(r.Version.Patch),
			},
		})
	}

	validator := &validatorpb.Validator{
		Id:                rawValidator.ID.String(),
		Address:           address,
		TendermintAddress: tmAddr,
		VotingPower:       rawValidator.VotingPower,

		Node: &validatorpb.Node{
			Id:         rawNode.ID.String(),
			EntityId:   rawNode.EntityID.String(),
			Expiration: rawNode.Expiration,

			P2P: &validatorpb.P2PInfo{
				Id:        rawNode.P2P.ID.String(),
				Addresses: p2pAddresses,
			},

			Consensus: &validatorpb.ConsensusInfo{
				Id:        rawNode.Consensus.ID.String(),
				Addresses: consensusAddresses,
			},

			Runtimes: runtimes,
			Roles:    uint32(rawNode.Roles),
		},
	}

	if rateNumerator != nil {
		rate := big.NewRat(rateNumerator.ToBigInt().Int64(), stakingApi.CommissionRateDenominator.ToBigInt().Int64())
		rate.Mul(rate, big.NewRat(100, 1))
		validator.Commission = rateNumerator.ToBigInt().Bytes()
	}

	return validator
}
