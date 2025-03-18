package recorder

import (
	"bufio"
	"io"
	"os"

	"github.com/koyo-os/log-saver/internal/config"
	"github.com/koyo-os/log-saver/pkg/logger"
)

type Recorder struct{
	OutputChan chan []byte
	logger *logger.Logger
	cfg *config.ConsumerConfig
	inputs io.Reader
}

func Init(cfg *config.ConsumerConfig, logger *logger.Logger, outputCh chan []byte) *Recorder {
	var input io.Reader

	switch cfg.Loginput{
	case "stdout":
		input = os.Stdout
	case "stderr":
		input = os.Stderr
	}

	return &Recorder{
		inputs: input,
		cfg: cfg,
		logger: logger,
		OutputChan: outputCh,
	}
}

func (r *Recorder) Run() {
	scanner := bufio.NewScanner(r.inputs)

	scanner.Split(bufio.ScanLines)

	for {
		if scanner.Scan() {
			line := scanner.Text()
			r.OutputChan <- []byte(line	)
		} else {
			break
		}
	}
}