package config

import (
	"flag"
	"metricsserviceGRPC/internal/constants"
	"metricsserviceGRPC/internal/logger"

	"github.com/caarlos0/env"
)

type ConfigServer struct {
	Port     string `env:"ADDRESS"`
	Interval int    `env:"STORE_INTERVAL"`
	FileName string `env:"FILE_STORAGE_PATH"`
	Restore  bool   `env:"RESTORE"`
	DSN      string `env:"DATABASE_DSN"`
	HashKey  string `env:"KEY"`
	Logger   *logger.Config
}

func InitServer() (*ConfigServer, error) {

	var flagRunAddr string
	var flagInterval int
	var flagFileName string
	var flagRestore bool
	var flagDSN string
	var flaghashkey string

	cfg := &ConfigServer{}
	_ = env.Parse(cfg)

	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&flagInterval, "i", 300, "Saved interval")
	flag.StringVar(&flagFileName, "f", "/tmp/metrics-db.json", "storage file")
	flag.BoolVar(&flagRestore, "r", true, "need of sviving")
	flag.StringVar(&flaghashkey, "k", "secretkey", "key for hash func")

	flag.StringVar(&flagDSN, "d", "", "connection to database")

	flag.Parse()

	if cfg.Port == "" {
		cfg.Port = flagRunAddr
	}
	if cfg.FileName == "" {
		cfg.FileName = flagFileName
	}
	if cfg.Interval == 0 {
		cfg.Interval = flagInterval
	}

	if !cfg.Restore {
		cfg.Restore = flagRestore
	}

	if cfg.DSN == "" {
		cfg.DSN = flagDSN
	}

	if cfg.HashKey == "" {
		cfg.HashKey = flaghashkey
	}

	cfglog := &logger.Config{
		LogLevel: constants.LogLevel,
		DevMode:  constants.DevMode,
		Type:     constants.Type,
	}

	cfg.Logger = cfglog

	return cfg, nil
}
