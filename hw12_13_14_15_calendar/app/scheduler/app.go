package scheduler

import (
	"context"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/app/scheduler/domain"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/config"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/messagequeue"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type Scheduler struct {
	eventStorage storage.EventStorage
	logger       logger.Logger
	config       *config.Config
	mq           messagequeue.MessageQueue
}

func NewScheduler(
	eventStorage storage.EventStorage,
	logger logger.Logger,
	config *config.Config,
	mq messagequeue.MessageQueue,
) *Scheduler {
	return &Scheduler{eventStorage: eventStorage, logger: logger, config: config, mq: mq}
}

func (a *Scheduler) Start(ctx context.Context) error {
	err := a.mq.Open()
	a.logger.Info("scheduler starting")
	if err != nil {
		return err
	}
	frequency, err := time.ParseDuration(a.config.SchedulerPeriod)
	if err != nil {
		frequency = time.Minute
	}

	ticker := time.NewTicker(frequency)
	go func() {
		for {
			select {
			case <-ctx.Done():
				a.mq.Close()
				return
			case <-ticker.C:
				err := a.notificationToQueue(frequency)
				if err != nil {
					a.logger.Error(err.Error())
				}

				_, err = a.removeOldEvents(frequency)
				if err != nil {
					a.logger.Error(err.Error())
				}
			}
		}
	}()

	return nil
}

func (a *Scheduler) notificationToQueue(frequency time.Duration) error {
	endDate := time.Now().Round(frequency)
	startDate := endDate.Add(-frequency)
	useCase := domain.NewUseCaseNotificationToQueue(a.eventStorage, a.logger, a.mq)

	return useCase.PutByPeriod(startDate, endDate)
}

func (a *Scheduler) removeOldEvents(frequency time.Duration) (int32, error) {
	year := time.Hour * 24 * 365

	dateForDel := time.Now().Round(frequency).Add(-year)

	return domain.NewUseCaseRemoveOldEvents(a.eventStorage).RemoveOlderThan(dateForDel)
}
