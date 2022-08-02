package event

import (
	"errors"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

var ErrNotFound = errors.New("Event not found")

type UseCaseFindEvent struct {
	storage storage.EventStorage
}

func NewUseCaseFindEvent(storage storage.EventStorage) *UseCaseFindEvent {
	return &UseCaseFindEvent{storage: storage}
}

func (u *UseCaseFindEvent) FindEvent(id int32, userID int32) (*Event, error) {
	eventDTO, err := u.storage.FindOneById(id)
	if err != nil {
		return nil, err
	}

	if eventDTO == nil {
		return nil, ErrNotFound
	}

	event := covert(eventDTO)

	rules := NewRules(u.storage)
	if err := rules.CheckCreateAccess(event, userID); err != nil {
		return nil, err
	}

	return event, nil
}

func (u *UseCaseFindEvent) FindEventsForDay(startDate time.Time, userID int32) ([]*Event, error) {
	endDate := startDate.AddDate(0, 0, 1)
	eventDTOs, err := u.storage.FindListByPeriod(startDate, endDate, userID)

	return convertList(eventDTOs), err
}

func (u *UseCaseFindEvent) FindEventsForWeek(startDate time.Time, userID int32) ([]*Event, error) {
	endDate := startDate.AddDate(0, 0, 7)
	eventDTOs, err := u.storage.FindListByPeriod(startDate, endDate, userID)

	return convertList(eventDTOs), err
}

func (u *UseCaseFindEvent) FindEventsForMonth(startDate time.Time, userID int32) ([]*Event, error) {
	endDate := startDate.AddDate(0, 1, 0)
	eventDTOs, err := u.storage.FindListByPeriod(startDate, endDate, userID)

	return convertList(eventDTOs), err
}

func convertList(eventDTOs []*storage.EventDTO) []*Event {
	events := make([]*Event, len(eventDTOs))
	for i := range eventDTOs {
		events[i] = covert(eventDTOs[i])
	}

	return events
}

func covert(e *storage.EventDTO) *Event {
	return &Event{
		id:                 e.ID(),
		title:              e.Title(),
		description:        e.Description(),
		userID:             e.UserID(),
		startDate:          e.StartDate(),
		endDate:            e.EndDate(),
		notificationBefore: e.NotificationBefore(),
	}
}
