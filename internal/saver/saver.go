package saver

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/bytedance/sonic"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/koyo-os/log-saver/internal/config"
	"github.com/koyo-os/log-saver/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Saver struct{
	client *elasticsearch.Client
	cfg *config.Config
	logger *logger.Logger
	ctx context.Context
}

const MAX_COUNTER = 50

func Init(cfg *config.Config, logger *logger.Logger) (*Saver, error) {
	c := elasticsearch.Config{
		Addresses: []string{
			cfg.Consumer.ELSurl,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 25 * time.Second)
	defer cancel()

	logger.Info("get elasticsearch client...")

	client, err := elasticsearch.NewClient(c)
	if err != nil{
		logger.Error("cant get els client", zapcore.Field{
			Key: "err",
			String: err.Error(),
		})
	}

	return &Saver{
		client: client,
		logger: logger,
		cfg: cfg,
		ctx: ctx,
	}, nil
}

var (
	count int
	body string
)

func (s *Saver) Save(b string) error {
	if count > MAX_COUNTER {
		doc := map[string]string{
			"date" : time.Now().Format("2006-01-02 15:04:05"),
			"logs" : body, 
		}

		var buf bytes.Buffer
		if err := sonic.ConfigDefault.NewEncoder(&buf).Encode(&doc);err != nil{
			return fmt.Errorf("cant marshal doc: %v", err)
		}

		return nil
	} else {
		count++
		body = fmt.Sprintf("%s%s", body, b)

		return nil
	}
}