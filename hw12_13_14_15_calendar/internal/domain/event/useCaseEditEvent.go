package event

import (
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type UseCaseEditEvent struct {
	storage storage.EventStorage
	rules   *Rules
}

func NewUseCaseEditEvent(storage storage.EventStorage) *UseCaseEditEvent {
	rules := NewRules(storage)

	return &UseCaseEditEvent{storage: storage, rules: rules}
}

func (u *UseCaseEditEvent) EditEvent(
	eventForUpdate *Event,
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
	authorUserID int,
) (bool, error) {
	if err := u.rules.CheckEditAccess(eventForUpdate, authorUserID); err != nil {
		return false, err
	}

	if err := u.rules.CheckDate(eventForUpdate); err != nil {
		return false, err
	}

	EventDTO := storage.NewEventDTO(
		eventForUpdate.Id(),
		title,
		description,
		eventForUpdate.UserID(),
		startDate,
		endDate,
		notificationBefore,
	)

	success, err := u.storage.Update(eventForUpdate.Id(), EventDTO)

	if success {
		eventForUpdate.title = title
		eventForUpdate.description = description
		eventForUpdate.startDate = startDate
		eventForUpdate.endDate = endDate
		eventForUpdate.notificationBefore = notificationBefore
	}

	return success, err
}
