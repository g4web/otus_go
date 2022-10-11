package domain

import (
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
	"github.com/jinzhu/copier"
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
	authorUserID int32,
) error {
	if err := u.rules.CheckEditAccess(eventForUpdate, authorUserID); err != nil {
		return err
	}

	eventForValidate := &Event{}
	copier.Copy(&eventForValidate, &eventForUpdate)

	eventForValidate.title = title
	eventForValidate.description = description
	eventForValidate.startDate = startDate
	eventForValidate.endDate = endDate
	eventForValidate.notificationBefore = notificationBefore

	if err := u.rules.CheckDate(eventForValidate); err != nil {
		return err
	}

	EventDTO := storage.NewEventDTO(
		eventForUpdate.ID(),
		title,
		description,
		eventForUpdate.UserID(),
		startDate,
		endDate,
		notificationBefore,
		false,
	)

	err := u.storage.Update(eventForUpdate.ID(), EventDTO)

	if err == nil {
		eventForUpdate.title = title
		eventForUpdate.description = description
		eventForUpdate.startDate = startDate
		eventForUpdate.endDate = endDate
		eventForUpdate.notificationBefore = notificationBefore
	}

	return err
}
