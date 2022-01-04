package storage

import (
	"github.com/NbNative/nbnative-server/drive/es"
	"net/http"
)

type EsStorage struct {
	client es.ElasticsearchClient
	Config es.ClientConfig
}

func (e EsStorage) GetSpan() {
	//TODO implement me
	panic("implement me")
}

func (e EsStorage) GetService() {
	//TODO implement me
	panic("implement me")
}

func NewEsClientStorage(config es.ClientConfig) (*EsStorage, error) {
	storage := &EsStorage{Config: config}
	client, err := es.NewElasticsearch7Client(config, &http.Transport{})
	storage.client = client
	return storage, err
}
