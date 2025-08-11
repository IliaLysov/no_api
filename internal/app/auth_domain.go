package app

import (
	"no_api/internal/auth/adapter/jwt"
	"no_api/internal/auth/adapter/kafka_producer"
	"no_api/internal/auth/adapter/postgres"
	"no_api/internal/auth/controller/http_router"
	"no_api/internal/auth/usecase"
)

func AuthDomain(d Dependencies) {
	authUseCase := usecase.New(
		postgres.New(d.Postgres.Pool),
		jwt.New(),
		d.Email,
		kafka_producer.New(d.KafkaWriter.Writer),
	)

	http_router.AuthRouter(d.RouterHTTP, authUseCase)
}
