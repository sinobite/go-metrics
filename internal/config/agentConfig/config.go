package agentConfig

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
		FlagRunEndpoint: "localhost:8089",
		ReportInterval:  10,
		PollInterval:    2,
	}

	return cfg
}

func (c EnvConfig) Parse() {
	c.parseEnvs()
	c.parseFlags()
}

func (c EnvConfig) parseEnvs() {
	err := env.Parse(&c)
	if err != nil {
		log.Fatal(err)
	}
}

func (c EnvConfig) parseFlags() {
	flag.StringVar(&c.FlagRunEndpoint, "a", "localhost:8089", "address and port to run server")
	flag.IntVar(&c.ReportInterval, "r", 10, "report Interval for metrics")
	flag.IntVar(&c.PollInterval, "p", 2, "pool Interval for metrics")
	flag.Parse()

}
