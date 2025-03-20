package saver

import (
	"bytes"
	"encoding/json"
	"fmt"

	"go.uber.org/zap/zapcore"
)

func (s *Saver) GetLogs(date string) (map[string]interface{}, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"date": date,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query);err != nil{
		return map[string]interface{}{}, fmt.Errorf("camt encode body: %v",err)
	}

	res, err := s.client.Search(
		s.client.Search.WithContext(s.ctx),
		s.client.Search.WithIndex(s.cfg.Consumer.IndexName),
		s.client.Search.WithBody(&buf),
		s.client.Search.WithPretty(),
	)

	if err != nil{
		return map[string]interface{}{}, fmt.Errorf("cant get resp: %v", err)
	}
	defer res.Body.Close()

	var result map[string]interface{}

	if err = json.NewDecoder(res.Body).Decode(&result);err != nil{
		s.logger.Error("cant decode resp", zapcore.Field{
			Key: "err",
			String: err.Error(),
		})
		return map[string]interface{}{}, fmt.Errorf("cant unmarshal resp: %v", err)
	}
	
	s.logger.Info("success reciew logs for", zapcore.Field{
		Key: "date",
		String: date,
	})

	return result, nil
}