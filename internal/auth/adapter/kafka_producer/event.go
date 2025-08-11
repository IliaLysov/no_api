package kafka_producer

import (
	"context"
	"fmt"
	"no_api/internal/auth/entity"

	"github.com/segmentio/kafka-go"
)

func (p *Producer) CreateEvent(ctx context.Context, e entity.CreateEvent) error {
	m := kafka.Message{
		Key:   []byte(e.ID),
		Value: []byte(e.Name),
	}

	err := p.writer.WriteMessages(ctx, m)
	if err != nil {
		return fmt.Errorf("p.writer.WriteMessages: %w", err)
	}

	return nil
}
