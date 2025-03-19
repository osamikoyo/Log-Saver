package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Consumer struct{
	KafkaUrl string `yaml:"kafka_url"`
	ELSurl string `yaml:"els_url"`
}

type Producer struct{
	KafkaUrl string `yaml:"kafka_url"`
	ELSurl string `yaml:"els_url"`
	Input string `yaml:"input"`
}

type Config struct{
	Topic string `yaml:"topic"`
	Consumer `yaml:"consumer"`
	Producer `yaml:"producer"`
}

func Init() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil{
		return nil, fmt.Errorf("cant open config file: %v", err)
	}

	body, err := io.ReadAll(file)
	if err != nil{
		return nil, fmt.Errorf("cant read file: %v", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(body, &cfg);err != nil{
		return nil, fmt.Errorf("cant unmarshal: %v", err)
	}

	return &cfg, nil
}