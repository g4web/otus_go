package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/app/calendar"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/config"
	servergrpc "github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/grpc"
	serverhttp "github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/http"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./config/config.env", "path to configuration file")
}

func main() {
	flag.Parse()

	configs, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	logg := logger.New(configs.LogLevel, configs.LogFile)
	defer logg.Close()

	eventStorage, err := getStorage(configs)
	if err != nil {
		logg.Error(err.Error())
	}

	calendarApp := calendar.New(logg, eventStorage)

	httpServer := serverhttp.NewServer(logg, calendarApp, configs)
	grpcServer := servergrpc.NewServer(logg, calendarApp, configs)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		logg.Info("running http server...")
		if err := httpServer.Start(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Error("Fail to run http server " + err.Error())
		}
	}()

	go func() {
		logg.Info("running grpc server...")
		if err := grpcServer.Start(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Error("fail to run grpc server " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	<-ctx.Done()

	if err := httpServer.Stop(ctx); err != nil {
		logg.Error(err.Error())
	}
	if err := grpcServer.Stop(ctx); err != nil {
		logg.Error(err.Error())
	}
}

func getStorage(c *config.Config) (storage.EventStorage, error) {
	switch c.StorageType {
	case "memory":
		return memorystorage.New(), nil
	case "postgres":
		return sqlstorage.New(c)
	default:
		return nil, errors.New("unsupported storage type")
	}
}
