package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Title string `json:"title"`
		Port  uint16 `json:"port"`
	}
}

func NewConfig(yamlFile string) (*Config, error) {
	f, err := os.Open(yamlFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
