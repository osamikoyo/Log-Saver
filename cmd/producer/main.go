package main

import (
	"github.com/koyo-os/log-saver/internal/config"
	"github.com/koyo-os/log-saver/internal/producer"
	"github.com/koyo-os/log-saver/internal/recorder"
	"github.com/koyo-os/log-saver/pkg/logger"
)

func main() {
	logger := logger.Init()

	cfg, err := config.Init()
	if err != nil{
		logger.Error(err.Error())
		return
	}

	var ch chan []byte

	producer, err := producer.Init(cfg, logger, ch)
	if err != nil{
		logger.Error(err.Error())
		return
	}

	recorder := recorder.Init(cfg, logger, ch)
	go producer.Listen()
	recorder.Run()
}