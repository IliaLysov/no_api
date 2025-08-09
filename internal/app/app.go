package app

import (
	"context"
	"fmt"
	"no_api/pkg/http_server"
	"no_api/pkg/postgres"
	"no_api/pkg/router"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

type Dependencies struct {
	RouterHTTP *chi.Mux
	Postgres   *postgres.Pool
}

func Run(ctx context.Context) (err error) {
	var deps Dependencies

	deps.Postgres, err = postgres.New(ctx)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}
	defer deps.Postgres.Close()

	deps.RouterHTTP = router.New()

	AuthDomain(deps)

	httpServer := http_server.New(deps.RouterHTTP, "8080")
	defer httpServer.Close()

	waiting(httpServer)

	return nil
}

func waiting(httpServer *http_server.Server) {
	fmt.Println("App started")
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case <-wait:
		fmt.Println("Received SIGTERM")
	case err := <-httpServer.Notify():
		fmt.Println("Received error:", err)
	}
	fmt.Println("Received exit signal")
}
