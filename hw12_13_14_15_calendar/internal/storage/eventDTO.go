package storage

import (
	"time"
)

type EventDTO struct {
	id                 int32
	title              string
	description        string
	userID             int32
	startDate          time.Time
	endDate            time.Time
	notificationBefore int32
	notificationIsSent bool
}

func NewEventDTO(
	id int32,
	title string,
	description string,
	userID int32,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
) *EventDTO {
	return &EventDTO{
		id:                 id,
		title:              title,
		description:        description,
		userID:             userID,
		startDate:          startDate,
		endDate:            endDate,
		notificationBefore: int32(notificationBefore.Round(time.Second).Seconds()),
	}
}

func (e *EventDTO) SetID(id int32) {
	e.id = id
}

func (e EventDTO) ID() int32 {
	return e.id
}

func (e EventDTO) Title() string {
	return e.title
}

func (e EventDTO) Description() string {
	return e.description
}

func (e EventDTO) UserID() int32 {
	return e.userID
}

func (e EventDTO) StartDate() time.Time {
	return e.startDate
}

func (e EventDTO) EndDate() time.Time {
	return e.endDate
}

func (e EventDTO) NotificationBefore() time.Duration {
	return time.Duration(float64(e.notificationBefore) * 1e9)
}

func (e *EventDTO) MarkNotificationAsSent() {
	e.notificationIsSent = true
}

func (e *EventDTO) NotificationIsSent() bool {
	return e.notificationIsSent
}
