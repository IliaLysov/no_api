package kafka_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"no_api/internal/auth/entity"

	"github.com/segmentio/kafka-go"
)

func (p *Producer) CreateEvent(ctx context.Context, e entity.CreateEvent) error {
	b, err := json.Marshal(e)

	if err != nil {
		return err
	}

	m := kafka.Message{
		Key:   []byte(e.ID),
		Value: b,
	}

	err = p.writer.WriteMessages(ctx, m)
	if err != nil {
		return fmt.Errorf("p.writer.WriteMessages: %w", err)
	}

	return nil
}
