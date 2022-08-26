package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/app/scheduler"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/config"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/messagequeue"
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

	mq := messagequeue.NewRabbitMQ(configs.MQAddr, configs.MQQueue, configs.MQHandlersCount, logg)
	app := scheduler.NewScheduler(eventStorage, logg, configs, mq)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logg.Error("waiting for the rabbit to start...")
	time.Sleep(time.Second * 10) // waiting for the rabbit to start

	err = app.Start(ctx)
	if err != nil {
		logg.Error(err.Error())
	}

	<-ctx.Done()
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
