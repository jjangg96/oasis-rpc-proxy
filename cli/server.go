package cli

import (
	"github.com/jjangg96/oasis-rpc-proxy/config"
	"github.com/jjangg96/oasis-rpc-proxy/server"
)

func startServer(cfg *config.Config) error {
	doc, err := initGenesis(cfg)
	if err != nil {
		return err
	}
	client, err := initClient(cfg)
	if err != nil {
		return err
	}
	defer client.Close()

	a := server.New(cfg, client, doc)
	if err := a.Start(cfg.ServerAddr, cfg.ServerPort); err != nil {
		return err
	}
	return nil
}
