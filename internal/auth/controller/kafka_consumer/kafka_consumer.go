package kafka_consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"no_api/internal/auth/entity"
	"no_api/pkg/kafka_reader"

	"github.com/rs/zerolog/log"
)

type Handler func(ctx context.Context, env entity.CreateEvent) error

func AuthConsumer(reader *kafka_reader.Reader) {
	ctx := context.Background()

	handlers := map[string]Handler{
		"user.created":   handleUserCreated,
		"user.logged_in": handleUserLoggedIn,
		"user.protected": handleProtected,
	}

	for {
		m, err := reader.FetchMessage(ctx)
		if err != nil {
			fmt.Println("kafka_consumer.AuthConsumer: reader.FetchMessage")
		}

		var raw map[string]json.RawMessage

		if err := json.Unmarshal(m.Value, &raw); err != nil {
			log.Error().Err(err).Msg("unmarshal envelope")

			_ = reader.CommitMessages(ctx, m)
			continue
		}

		var t string
		if err := json.Unmarshal(raw["type"], &t); err != nil {
			log.Error().Err(err).Msg("no type in event")
			_ = reader.CommitMessages(ctx, m)
			continue
		}

		var env entity.CreateEvent
		if err := json.Unmarshal(m.Value, &env); err != nil {
			log.Error().Err(err).Msg("unmarshal envelope full")
			_ = reader.CommitMessages(ctx, m)
			continue
		}
		h, ok := handlers[t]

		if !ok {
			log.Warn().Str("type", t).Msg("no handler for event type")
			_ = reader.CommitMessages(ctx, m)
			continue
		}

		if err := h(ctx, env); err != nil {
			log.Error().Err(err).Str("type", t).Msg("handler error")
			continue
		}

		if err = reader.CommitMessages(ctx, m); err != nil {
			log.Error().Err(err).Msg("kafka_consumer.AuthConsumer: reader.CommitMessages")
		}
	}
}

func handleUserCreated(ctx context.Context, env entity.CreateEvent) error {
	var p entity.UserCreated
	b, _ := json.Marshal(env.Payload)
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}
	log.Info().Str("Created user with email:", p.Email).Msg("kafka_consumer.AuthConsumer: handleUserCreated")
	return nil
}

func handleUserLoggedIn(ctx context.Context, env entity.CreateEvent) error {
	var p entity.UserLoggedIn
	b, _ := json.Marshal(env.Payload)
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}
	log.Info().Str("Logged in user with email:", p.Email).Str("User ip:", p.IP).Msg("kafka_consumer.AuthConsumer: handleUserLoggedIn")
	return nil
}

func handleProtected(ctx context.Context, env entity.CreateEvent) error {
	var p entity.UserProtected
	b, _ := json.Marshal(env.Payload)
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}
	log.Info().Str("Logged in user with id:", p.ID).Str("User id:", p.ID).Msg("kafka_consumer.AuthConsumer: handleProtected")
	return nil
}
