package notify

import (
	"context"
	"fmt"
	"no_api/internal/auth/controller/kafka_consumer"
	"no_api/pkg/kafka_reader"
)

type Dependencies struct {
	KafkaReader *kafka_reader.Reader
}

func Run(ctx context.Context) (err error) {
	var deps Dependencies

	deps.KafkaReader, err = kafka_reader.New()
	if err != nil {
		return fmt.Errorf("kafka_reader.New: %w", err)
	}
	defer deps.KafkaReader.Close()

	kafka_consumer.AuthConsumer(deps.KafkaReader)

	return nil
}
