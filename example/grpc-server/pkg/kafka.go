package pkg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func Consume(address string) {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{address},
		Topic:     "my-topic",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
