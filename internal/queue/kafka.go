package queue

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/segmentio/kafka-go"
)

func StartReader(ctx context.Context, addrs []string, topic string) {
	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: addrs,
			Topic: topic,
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