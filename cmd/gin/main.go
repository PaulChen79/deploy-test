package main

import (
	route "deploy-test/internal/gin/route"
	"deploy-test/internal/logger"
	"deploy-test/internal/provider"

	"github.com/jpillora/overseer"
	"go.uber.org/zap"
)

func main() {
	config := provider.NewConfig()
	overseer.Run(overseer.Config{
		Program:   prog,
		Addresses: []string{":" + config.Gin.Port},
		Debug:     false, // display log of overseer actions
	})
}

func prog(state overseer.State) {

	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	config := provider.NewConfig()

	router := route.SetupRouter(config)
	router.RunListener(state.Listener)
}
