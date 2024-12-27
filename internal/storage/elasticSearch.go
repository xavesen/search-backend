package storage

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearchClient struct {
	Client 	*elasticsearch.TypedClient
}

func NewElasticSearchClient(addr []string, apiKey string) (*ElasticSearchClient, error) {
	// As it is an pet project addr will be http so no certs are considered
	cfg := elasticsearch.Config{
        Addresses: addr,
		APIKey: apiKey,
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Errorf("Error creating elastic search client: %s", err)
		return nil, err
	}

	err = es.BaseClient.DiscoverNodes()
	if err != nil {
		log.Errorf("Error connecting to elastic search on %s: %s", strings.Join(addr, ", "), err)
		return nil, err
	}

	return &ElasticSearchClient{Client: es}, nil
}