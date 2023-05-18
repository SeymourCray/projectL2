package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port string `yaml:"port"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	// Open config file
	file, err := os.Open("../config.yml")
	if err != nil {
		return nil, err
	}

	defer file.Close()
	// Init new YAML decode
	d := yaml.NewDecoder(file)
	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
