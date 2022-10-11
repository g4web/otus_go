package domain

import (
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/messagequeue"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type UseCasePutNotificationToQueue struct {
	storage storage.EventStorage
	logger  logger.Logger
	mq      messagequeue.MessageQueue
}

func NewUseCaseNotificationToQueue(
	storage storage.EventStorage,
	logger logger.Logger,
	mq messagequeue.MessageQueue,
) *UseCasePutNotificationToQueue {
	return &UseCasePutNotificationToQueue{storage: storage, logger: logger, mq: mq}
}

func (u *UseCasePutNotificationToQueue) PutByPeriod(startDate time.Time, endDate time.Time) error {
	eventDTOs, err := u.storage.FindNotificationByPeriod(startDate, endDate)
	if err != nil {
		return err
	}

	for _, event := range eventDTOs {
		eventMq := u.convertEvent(event)
		if err := u.mq.Publish(eventMq); err != nil {
			u.logger.Error("fail publish event message")
		}
	}

	return nil
}

func (u *UseCasePutNotificationToQueue) convertEvent(event *storage.EventDTO) *messagequeue.EventDTO {
	return &messagequeue.EventDTO{
		ID:          event.ID(),
		Title:       event.Title(),
		Description: event.Description(),
		UserID:      event.UserID(),
		StartDate:   event.StartDate(),
		EndDate:     event.EndDate(),
	}
}
