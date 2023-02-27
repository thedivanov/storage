package main

import (
	"github.com/caarlos0/env"
)

type config struct {
	MemcachedAddr string `env:"MEMCACHED_ADDR,required"`
	GRPCAddr      string `env:"GRPC_ADDR,required"`
}

func parseConfig() (*config, error) {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
