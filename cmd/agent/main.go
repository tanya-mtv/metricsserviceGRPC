package main

import (
	"metricsserviceGRPC/internal/agent"
	"metricsserviceGRPC/internal/config"
	"metricsserviceGRPC/internal/logger"
)

func main() {
	cfg, err := config.InitAgent()
	if err != nil {

		panic("error initialazing config")
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()

	ag := agent.NewAgent(cfg, appLogger)

	if err := ag.Run(); err != nil {
		appLogger.Fatal(err)
	}

}
