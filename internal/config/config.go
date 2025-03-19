package config

type Consumer struct{
	KafkaUrl string `yaml:"kafka_url"`
}

type Producer struct{
	KafkaUrl string `yaml:"kafka_url"`
	ELSurl string `yaml:"els_url"`
	Input string `yaml:"input"`
}

type Config struct{
	Consumer `yaml:"consumer"`
	Producer `yaml:"producer"`
}