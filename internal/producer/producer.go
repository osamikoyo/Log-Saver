package producer

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/koyo-os/log-saver/internal/config"
	"github.com/koyo-os/log-saver/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Producer struct{
	logger *logger.Logger
	producer sarama.SyncProducer
	cfg *config.
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

func (p *Producer) Listen() {
	for {
		msg := &sarama.ProducerMessage{
			Topic: p.cfg.Topic,
			Value: sarama.StringEncoder(<- p.outputCh),	
		}

		part, offset, err := p.producer.SendMessage(msg)
		if err != nil{
			p.logger.Error("cant send message", zapcore.Field{
				Key: "err",
				String: err.Error(),
			})
		}

		p.logger.Info("send message with", zapcore.Field{
			Key: "offset",
			Integer: offset,
		},
		zapcore.Field{
			Key: "part",
			Integer: int64(part),
		})
	}
}