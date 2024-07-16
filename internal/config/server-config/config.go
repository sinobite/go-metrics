package server_config

import (
	"flag"
	"os"
)

type EnvConfig struct {
	FlagRunEndpoint string `env:"ADDRESS"`
}

func New() EnvConfig {
	cfg := EnvConfig{FlagRunEndpoint: "localhost:8080"}

	cfg.parseFlags()

	return cfg
}

func (c *EnvConfig) parseFlags() {
	flag.StringVar(&c.FlagRunEndpoint, "a", "localhost:8080", "address and port to run server")
	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		c.FlagRunEndpoint = envRunAddr
	}
}
