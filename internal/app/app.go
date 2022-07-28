package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chiefcake/apod/clients/apod"
	"github.com/chiefcake/apod/internal/config"
	"github.com/chiefcake/apod/internal/handler"
	"github.com/chiefcake/apod/internal/server"
	"github.com/chiefcake/apod/internal/service"
	"github.com/chiefcake/apod/internal/storage"
	"github.com/chiefcake/apod/internal/storage/postgres"
	"github.com/chiefcake/apod/internal/worker"
)

const timeout = 3 * time.Second

// Run starts the app gateway with provided config values.
func RunWithConfig(ctx context.Context, cfg *config.Config) {
	pg, err := storage.NewPostgres(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer pg.Close()

	apodClient, err := apod.NewClient(cfg.APIKey)
	if err != nil {
		log.Println(err)
		return
	}

	pictureRepository := postgres.NewPictureRepository(pg)
	pictureService := service.NewPicture(apodClient, pictureRepository)
	pictureHandler := handler.NewPicture(pictureService)

	worker := worker.New(pictureService)
	defer worker.Stop()

	go worker.Run(ctx)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	server := server.New(cfg, pictureHandler)

	log.Printf("Server is running on [%s:%s]...\n", cfg.ServerHost, cfg.ServerPort)

	go func() {
		err = server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	log.Println("Shutting down server...")

	err = server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
		return
	}
}

// Run parses the config values and starts the app gateway.
func Run(ctx context.Context) {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	RunWithConfig(ctx, cfg)
}
