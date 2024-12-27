package main

import (
	"context"
	"os"

	"github.com/xavesen/search-backend/internal/config"
	"github.com/xavesen/search-backend/internal/queue"
	"github.com/xavesen/search-backend/internal/storage"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		os.Exit(1)
	}

	_, err = storage.NewElasticSearchClient(config.ElasticSearchURLs, config.ElasticSearchKey)
	if err != nil {
		os.Exit(1)
	}
	
	queue.StartReader(context.TODO(), config.KafkaAddrs, config.KafkaTopic, config.KafkaGroupId, config.KafkaCommitInterval)
}