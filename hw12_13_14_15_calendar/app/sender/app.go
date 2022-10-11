package sender

import (
	"context"
	"fmt"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/config"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/messagequeue"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type Sender struct {
	eventStorage storage.EventStorage
	logger       logger.Logger
	config       *config.Config
	mq           messagequeue.MessageQueue
}

func NewSender(
	eventStorage storage.EventStorage,
	logger logger.Logger,
	config *config.Config, mq messagequeue.MessageQueue,
) *Sender {
	return &Sender{eventStorage: eventStorage, logger: logger, config: config, mq: mq}
}

func (s *Sender) Run(ctx context.Context) error {
	s.logger.Info("sender starting")
	if err := s.mq.Open(); err != nil {
		return err
	}

	return s.mq.Consume(ctx, s.eventHandle)
}

func (s *Sender) eventHandle(ctx context.Context, eventChannel <-chan *messagequeue.EventDTO) {
	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-eventChannel:
			if !ok {
				return
			}

			msg := fmt.Sprintf("received message from a queue: %+v", event)
			s.logger.Info(msg)

			if err := s.eventStorage.MarkNotificationAsSent(event.ID); err != nil {
				s.logger.Error("notify failed: " + err.Error())
			}
		}
	}
}

func (s *Sender) Shutdown() {
	s.logger.Info("sender is shutting down")
	s.mq.Close()
}
