package saver

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (s *Saver) GetLogs(date string) (string, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"date": date,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query);err != nil{
		return "", fmt.Errorf("camt encode body: %v",err)
	}

	res, err := s.client.Search(
		s.client.Search.WithContext(s.ctx),
	)
}