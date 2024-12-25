package main

import (
	"context"

	"github.com/xavesen/search-backend/internal/queue"
)

func main() {
	queue.StartReader(context.TODO(), []string{"localhost:9092"}, "test", "some-group-id", 5)
}