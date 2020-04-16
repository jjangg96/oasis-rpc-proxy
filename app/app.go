package app

import (
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/config"
	"github.com/figment-networks/oasis-rpc-proxy/utils/log"
	"github.com/gin-gonic/gin"
	genesisFile "github.com/oasislabs/oasis-core/go/genesis/file"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrl()

	log.Info("initializing Oasis genesis document")
	if err := initGenesis(); err != nil {
		panic(err)
	}

	port := config.AppPort()
	log.Info(fmt.Sprintf("Starting server at port %s...", port))
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}

func initGenesis() error {
	genesis, err := genesisFile.NewFileProvider("genesis.json")
	if err != nil {
		log.Error("failed to load genesis file", err)
		return err
	}

	// Retrieve the genesis document and use it to configure the ChainID for
	// signature domain separation. We do this as early as possible.
	doc, err := genesis.GetGenesisDocument()
	if err != nil {
		log.Error("failed to retrieve genesis document", err)
		return err
	}

	fmt.Printf("Chain context: '%v'\n\n", doc.ChainContext())
	fmt.Printf("Chain ID: '%v'\n\n", doc.Hash().String())
	var some []byte
	hash := doc.Hash()
	err2 := hash.UnmarshalBinary(some)
	if err2 != nil {

	}
	fmt.Printf("Chain ID: '%v'\n\n", some)

	doc.SetChainContext()

	return nil
}