package kafka_consumer

import (
	"context"
	"fmt"
	"no_api/pkg/kafka_reader"

	"github.com/rs/zerolog/log"
)

func AuthConsumer(reader *kafka_reader.Reader) {
	ctx := context.Background()

	for {
		m, err := reader.FetchMessage(ctx)
		if err != nil {
			fmt.Println("kafka_consumer.AuthConsumer: reader.FetchMessage")
		}
		log.Info().
			Str("topic", m.Topic).
			Int("partition", m.Partition).
			Int64("offset", m.Offset).
			Str("key", string(m.Key)).
			Str("value", string(m.Value)).
			Msg("kafka_consumer.AuthConsumer: reader.FetchMessage")
		if err = reader.CommitMessages(ctx, m); err != nil {
			log.Error().Err(err).Msg("kafka_consumer.AuthConsumer: reader.CommitMessages")
		}
	}
}
