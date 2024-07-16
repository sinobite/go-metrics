package agentconfig

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
)

type EnvConfig struct {
	FlagRunEndpoint string `env:"ADDRESS"`
	ReportInterval  int    `env:"REPORT_INTERVAL"`
	PollInterval    int    `env:"POLL_INTERVAL"`
}

func New() EnvConfig {
	cfg := EnvConfig{
		FlagRunEndpoint: "localhost:8080",
		ReportInterval:  10,
		PollInterval:    2,
	}

	return cfg
}

func (c *EnvConfig) Parse() {
	c.parseEnvs()
	c.parseFlags()
}

func (c *EnvConfig) parseEnvs() {
	err := env.Parse(c)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *EnvConfig) parseFlags() {
	flag.StringVar(&c.FlagRunEndpoint, "a", c.FlagRunEndpoint, "address and port to run server")
	flag.IntVar(&c.ReportInterval, "r", c.ReportInterval, "report Interval for metrics")
	flag.IntVar(&c.PollInterval, "p", c.PollInterval, "pool Interval for metrics")
	flag.Parse()

}
