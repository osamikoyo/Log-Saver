package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type ProducerConfig struct{
	KafkaUrl string `yaml:"kafka_url"`
	Input string `yaml:"input"`
	ELSurl string `yaml:"els_url"`
	Topic string `yaml:"topic"`
}

func InitPC() (*ProducerConfig,error) {
	file, err := os.Open("config.yaml")
	if err != nil{
		return nil,fmt.Errorf("cant open config file: %v",err)
	}

	body, err := io.ReadAll(file)
	if err != nil{
		return nil, fmt.Errorf("cant read body: %v",err)
	}

	var cfg ProducerConfig

	if err = yaml.Unmarshal(body, &cfg);err != nil{
		return nil, fmt.Errorf("cant unmarshal config body: %v",err)
	}

	return &cfg, nil
}