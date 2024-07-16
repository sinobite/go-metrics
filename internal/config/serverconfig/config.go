package serverconfig

import (
	"flag"
	"os"
)

type EnvConfig struct {
	FlagRunEndpoint string `env:"ADDRESS"`
}

func New() EnvConfig {
	cfg := EnvConfig{FlagRunEndpoint: "localhost:8080"}

	return cfg
}

func (c *EnvConfig) Parse() {
	c.parseEnvs()
	c.parseFlags()
}

func (c *EnvConfig) parseEnvs() {
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		c.FlagRunEndpoint = envRunAddr
	}
}

func (c *EnvConfig) parseFlags() {
	flag.StringVar(&c.FlagRunEndpoint, "a", "localhost:8080", "address and port to run server")
	flag.Parse()
}
