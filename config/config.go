// config.go
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
	} `yaml:"database"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
