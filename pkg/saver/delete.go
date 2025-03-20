package saver

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.uber.org/zap/zapcore"
)

func (s *Saver) Delete(id string) error {
	req := esapi.DeleteRequest{
		Index: s.cfg.Consumer.IndexName,
		DocumentID: id,
	}

	res, err := req.Do(s.ctx, s.client)
	if err != nil{
		 s.logger.Error("cant do req", zapcore.Field{
			Key: "err",
			String: err.Error(),
		 })
		 return err
	}
	defer res.Body.Close()

	s.logger.Info("success request", zapcore.Field{
		Key: "status",
		Integer: int64(res.StatusCode),
	})

}