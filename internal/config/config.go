package config

import (
	"diplom/pkg/logging"
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
)

type Config struct {
	IsDebug bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
	Storage `yaml:"storage"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func GetConfig() *Config {
	cfg := &Config{}

	logger := logging.GetLogger()
	logger.Info("read application configuration")

	loader := aconfig.LoaderFor(cfg, aconfig.Config{
		Files: []string{"config.yaml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
		},
	})

	err := loader.Load()
	if err != nil {
		panic(err)
	}

	return cfg
}
