package producer

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/koyo-os/log-saver/internal/config"
	"github.com/koyo-os/log-saver/pkg/logger"
)

type Producer struct{
	logger *logger.Logger
	producer sarama.SyncProducer
	cfg *config.ProducerConfig
	outputCh chan []byte
}

func Init(cfg *config.ProducerConfig, logger *logger.Logger, ch chan []byte) (*Producer, error) {
	configSarama := sarama.NewConfig()
	configSarama.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(
		[]string{cfg.KafkaUrl},
		configSarama,
	)
	defer producer.Close()

	if err != nil{
		return nil, fmt.Errorf("cant get producer: %v",err)
	}

	return &Producer{
		logger: logger,
		cfg: cfg,
		outputCh: ch,
		producer: producer,
	}, nil
}