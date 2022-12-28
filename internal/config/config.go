package config

import (
	"github.com/cristalhq/aconfig"
)

type Postgres struct {
	Host        string `env:"HOST"`
	Port        int    `env:"PORT"`
	Username    string `env:"USERNAME"`
	Password    string `env:"PASSWORD"`
	Database    string `env:"DATABASE"`
	SSLMode     string `env:"SSL_MODE"`
	SSLCertPath string `env:"SSL_CERT_PATH"`
}

type Config struct {
	Debug    bool     `env:"DEBUG"`
	Address  string   `env:"ADDRESS"`
	Postgres Postgres `env:"POSTGRES"`
}

func Load() (*Config, error) {
	cfg := &Config{}

	errLoad := aconfig.LoaderFor(cfg, aconfig.Config{
		EnvPrefix: "DIPLOM",
	}).Load()

	if errLoad != nil {
		return nil, errLoad
	}

	return cfg, nil
}
