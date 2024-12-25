package queue

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/segmentio/kafka-go"
)

func StartReader(ctx context.Context, addrs []string, topic string, groupId string, commitInterval int) {
	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: addrs,
			Topic: topic,
			GroupID: groupId,
			CommitInterval: time.Duration(commitInterval)*time.Second,
		},
	)

	defer reader.Close()

	for {
		document, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Error(err)
			break
		}
		log.Info(string(document.Value))
	}

	reader.Close()
}