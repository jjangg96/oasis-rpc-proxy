package mapper

import (
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/oasisprotocol/oasis-core/go/registry/api"
)

func RegistryToPb(rawRegistry api.Genesis) *statepb.Registry {
	// Gas Costs
	gasCosts := map[string]uint64{}
	for op, gas := range rawRegistry.Parameters.GasCosts {
		gasCosts[string(op)] = uint64(gas)
	}

	// Entities
	var entities []*statepb.Entity
	for _, rawEntity := range rawRegistry.Entities {
		entities = append(entities, &statepb.Entity{
			PublicKey: rawEntity.Signature.PublicKey.String(),
		})
	}

	return &statepb.Registry{
		Parameters: &statepb.RegistryParameters{
			DebugAllowUnroutableAddresses:          rawRegistry.Parameters.DebugAllowUnroutableAddresses,
			DebugAllowTestRuntimes:                 rawRegistry.Parameters.DebugAllowTestRuntimes,
			DebugBypassStake:                       rawRegistry.Parameters.DebugBypassStake,
			DisableRuntimeRegistration:             rawRegistry.Parameters.DisableRuntimeRegistration,
			DisableKeyManagerRuntimeRegistration:   rawRegistry.Parameters.DisableKeyManagerRuntimeRegistration,
			GasCosts:                               gasCosts,
			MaxNodeExpiration:                      rawRegistry.Parameters.MaxNodeExpiration,
		},
		Entities: entities,
	}
}
