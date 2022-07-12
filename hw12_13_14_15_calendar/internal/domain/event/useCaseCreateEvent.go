package event

import (
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type UseCaseCreateEvent struct {
	storage storage.EventStorage
	rules   *Rules
}

func NewUseCaseCreateEvent(storage storage.EventStorage) *UseCaseCreateEvent {
	rules := NewRules(storage)

	return &UseCaseCreateEvent{storage: storage, rules: rules}
}

func (u *UseCaseCreateEvent) CreateEvent(
	title string,
	description string,
	userID int,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
	authorUserID int,
) error {
	newEvent := NewEvent(
		title,
		description,
		userID,
		startDate,
		endDate,
		notificationBefore,
	)

	if err := u.rules.CheckCreateAccess(newEvent, authorUserID); err != nil {
		return err
	}

	if err := u.rules.CheckDate(newEvent); err != nil {
		return err
	}

	eventDTO := storage.NewEventDTO(
		0,
		newEvent.Title(),
		newEvent.Description(),
		newEvent.UserID(),
		newEvent.StartDate(),
		newEvent.EndDate(),
		newEvent.NotificationBefore(),
	)

	err := u.storage.Insert(eventDTO)
	if err != nil {
		return err
	}

	return err
}
