package agentconfig

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
)

type EnvConfig struct {
	FlagRunEndpoint  string `env:"ADDRESS"`
	ReportInterval   int    `env:"REPORT_INTERVAL"`
	PollInterval     int    `env:"POLL_INTERVAL"`
	AgentRunEndpoint string `env:"AGENT_ADDRESS"`
}

func New() EnvConfig {
	cfg := EnvConfig{
		FlagRunEndpoint:  "localhost:8080",
		ReportInterval:   10,
		PollInterval:     2,
		AgentRunEndpoint: "localhost:8089",
	}

	return cfg
}

func (c *EnvConfig) Parse() {
	err := c.parseEnvs()
	if err != nil {
		log.Fatalf("cant parse env variables: %s", err)
	}
	c.parseFlags()
}

func (c *EnvConfig) parseEnvs() error {
	err := env.Parse(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *EnvConfig) parseFlags() {
	flag.StringVar(&c.FlagRunEndpoint, "a", c.FlagRunEndpoint, "address and port to run server")
	flag.IntVar(&c.ReportInterval, "r", c.ReportInterval, "report Interval for metrics")
	flag.IntVar(&c.PollInterval, "p", c.PollInterval, "pool Interval for metrics")
	flag.StringVar(&c.AgentRunEndpoint, "aa", c.AgentRunEndpoint, "address and port to run agent")
	flag.Parse()

}
