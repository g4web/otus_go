package event

import (
	"time"
)

type Event struct {
	id                 int
	title              string
	description        string
	userID             int
	startDate          time.Time
	endDate            time.Time
	notificationBefore time.Duration
}

func NewEvent(
	title string,
	description string,
	userID int,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
) *Event {
	return &Event{
		title:              title,
		description:        description,
		userID:             userID,
		startDate:          startDate,
		endDate:            endDate,
		notificationBefore: notificationBefore,
	}
}

func (e *Event) Id() int {
	return e.id
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) Description() string {
	return e.description
}

func (e *Event) UserID() int {
	return e.userID
}

func (e *Event) StartDate() time.Time {
	return e.startDate
}

func (e *Event) EndDate() time.Time {
	return e.endDate
}

func (e *Event) NotificationDate() time.Time {
	return e.startDate.Add(e.notificationBefore)
}

func (e *Event) NotificationBefore() time.Duration {
	return e.notificationBefore
}
