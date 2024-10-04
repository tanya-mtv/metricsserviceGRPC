package main

import (
	"metricsserviceGRPC/internal/config"
	"metricsserviceGRPC/internal/logger"
	"metricsserviceGRPC/internal/server"
)

func main() {

	cfg, err := config.InitServer()
	if err != nil {

		panic("error initialazing config")
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()

	srv := server.NewServer(cfg, appLogger)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
