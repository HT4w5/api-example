package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HttpAddr string `yaml:"http_addr"`
}

func New() *Config {
	return &Config{
		HttpAddr: "0.0.0.0:80",
	}
}

func (cfg *Config) Load() {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Panicf("Failed to open config file: %v\n", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Panicf("Failed to read config file: %v\n", err)
	}

	var tmp Config
	err = yaml.Unmarshal(data, &tmp)
	if err != nil {
		log.Panicf("Failed to parse config file: %v\n", err)
	}

	// Check for empty config.
	if tmp.HttpAddr != "" {
		cfg.HttpAddr = tmp.HttpAddr
	}

}
