package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/configs"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage/sql"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/app"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.env", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := configs.NewConfig(configFile)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	logg := logger.New(config.LogLevel, config.LogFile)
	defer logg.Close()

	eventStorage, err := getStorage(config)
	if err != nil {
		logg.Error(err.Error())
	}

	calendar := app.New(logg, eventStorage)

	server := internalhttp.NewServer(logg, calendar, config)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}

func getStorage(c *configs.Config) (storage.EventStorage, error) {
	switch c.StorageType {
	case "memory":
		return memorystorage.New(), nil
	case "postgres":
		return sqlstorage.New(c)
	default:
		return nil, errors.New("Unsupported storage type")
	}
}
