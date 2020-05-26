package cli

import (
	"github.com/figment-networks/oasis-rpc-proxy/config"
	"github.com/figment-networks/oasis-rpc-proxy/grpc"
)

func startGrpc(cfg *config.Config) error {
	doc, err := initGenesis(cfg)
	if err != nil {
		return err
	}
	client, err := initClient(cfg)
	if err != nil {
		return err
	}
	defer client.Close()

	a := grpc.New(doc, client)
	if err := a.Start(cfg.ServerAddr, cfg.ServerPort); err != nil {
		return err
	}
	return nil
}
