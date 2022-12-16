package config

import (
	"github.com/cristalhq/aconfig"
)

type Config struct {
	IsDebug bool
	Listen  struct {
		Type   string `yaml:"type"`
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
}

func GetConfig() aconfig.Loader {
	var cfg Config

	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		Files: []string{"/config.yaml"},
	})

	err := loader.Load()
	if err != nil {
		panic(err)
	}

	return *loader
}

