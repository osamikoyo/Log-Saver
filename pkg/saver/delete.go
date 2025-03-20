package saver

import "github.com/elastic/go-elasticsearch/v8/esapi"

func (s *Saver) Delete(id string) error {
	req := esapi.DeleteRequest{
		Index: s.cfg.Consumer.IndexName,
		DocumentID: id,
	}

	
}