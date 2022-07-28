package main

import (
	"context"

	"github.com/chiefcake/apod/internal/app"
)

// @title APOD API
// @version 1.0

// @host localhost:8080
func main() {
	ctx := context.Background()

	app.Run(ctx)
}
