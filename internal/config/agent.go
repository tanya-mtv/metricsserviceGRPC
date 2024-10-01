package config

import (
	"flag"
	"metricsserviceGRPC/internal/constants"
	"metricsserviceGRPC/internal/logger"

	"github.com/caarlos0/env"
)

type ConfigAgent struct {
	Port           string `env:"ADDRESS"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	PollInterval   int    `env:"POLL_INTERVAL"`
	RateLimit      int    `env:"RATE_LIMIT"`
	HashKey        string `env:"KEY"`
	Logger         *logger.Config
}

func InitAgent() (*ConfigAgent, error) {
	var flagRunAddr string
	var pollInterval int
	var reportInterval int
	var flagRateLimit int
	var flaghashkey string

	cfg := &ConfigAgent{}
	_ = env.Parse(cfg)

	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&reportInterval, "r", 10, "report interval in seconds")
	flag.IntVar(&flagRateLimit, "l", 3, "rate limit")
	flag.IntVar(&pollInterval, "p", 2, "poll interval in seconds")
	flag.StringVar(&flaghashkey, "k", "secretkey", "key for hash func")

	flag.Parse()

	if cfg.PollInterval == 0 {
		cfg.PollInterval = pollInterval
	}

	if cfg.RateLimit == 0 {
		cfg.RateLimit = flagRateLimit
	}

	if cfg.ReportInterval == 0 {
		cfg.ReportInterval = reportInterval
	}

	if cfg.Port == "" {
		cfg.Port = flagRunAddr
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
