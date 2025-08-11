package http_router

import (
	"context"
	"fmt"
	"net/http"
	ver1 "no_api/internal/auth/controller/http_router/v1"
	"no_api/internal/auth/entity"
	"no_api/internal/auth/usecase"
	"strings"

	"github.com/go-chi/chi/v5"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(u *usecase.UseCase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "missing or invalid header", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			sub, err := u.JWT.Verify(tokenStr)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), userIDKey, sub)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthRouter(r *chi.Mux, uc *usecase.UseCase) {
	r.Route("/auth", func(r chi.Router) {
		v1 := ver1.New(uc)

		r.Route("/v1", func(r chi.Router) {
			r.Post("/signup", v1.CreateUser)
			r.Post("/login", v1.Login)
			r.Get("/tables", v1.GetTables)

			r.With(AuthMiddleware(uc)).Get("/protected", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				id, ok := r.Context().Value(userIDKey).(string)
				fmt.Println("id", id)

				event := entity.CreateEvent{
					ID:   id,
					Name: fmt.Sprintf("protected route by %s", id),
				}

				err := uc.Kafka.CreateEvent(r.Context(), event)
				if err != nil {
					fmt.Println("protected create event error", err)
				}

				if !ok {
					http.Error(w, "error to get id from token", http.StatusUnauthorized)
					return
				}
				w.Write([]byte("success\n"))
			}))
		})
	})
}
