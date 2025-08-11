package kafka_writer

import (
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

type Writer struct {
	*kafka.Writer
}

func New() (*Writer, error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_BROKERS")),
		Topic:    "auth-topic",
		Balancer: &kafka.LeastBytes{},
	}

	return &Writer{Writer: w}, nil
}

func (w *Writer) Close() {
	err := w.Writer.Close()
	if err != nil {
		fmt.Println("kafka_writer.Close")
	}
	fmt.Println("Kafka writer closed")
}
