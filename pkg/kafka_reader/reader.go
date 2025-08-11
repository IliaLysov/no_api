package kafka_reader

import (
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

type Reader struct {
	*kafka.Reader
}

func New() (*Reader, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_BROKERS")},
		GroupID:  "auth-group",
		Topic:    "auth-topic",
		MaxBytes: 10e6,
	})

	return &Reader{Reader: r}, nil
}

func (r *Reader) Close() {
	err := r.Reader.Close()
	if err != nil {
		fmt.Println("kafka_reader.Close")
	}
	fmt.Println("Kafka reader closed")
}
