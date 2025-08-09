package main

import (
	"context"
	"fmt"
	"no_api/config"
	"no_api/internal/app"
)

func main() {
	ctx := context.Background()

	err := config.New()
	if err != nil {
		fmt.Println("Config init error:", err)
	}

	err = app.Run(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println("App stopped")
}
