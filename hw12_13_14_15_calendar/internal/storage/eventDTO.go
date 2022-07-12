package storage

import "time"

type EventDTO struct {
	id                 int
	title              string
	description        string
	userID             int
	startDate          time.Time
	endDate            time.Time
	notificationBefore int
}

func NewEventDTO(id int, title string, description string, userID int, startDate time.Time, endDate time.Time, notificationBefore time.Duration) *EventDTO {
	return &EventDTO{
		id:                 id,
		title:              title,
		description:        description,
		userID:             userID,
		startDate:          startDate,
		endDate:            endDate,
		notificationBefore: int(notificationBefore.Round(time.Second).Seconds()),
	}
}

func (e *EventDTO) SetId(id int) {
	e.id = id
}

func (e EventDTO) ID() int {
	return e.id
}

func (e EventDTO) Title() string {
	return e.title
}

func (e EventDTO) Description() string {
	return e.description
}

func (e EventDTO) UserID() int {
	return e.userID
}

func (e EventDTO) StartDate() time.Time {
	return e.startDate
}

func (e EventDTO) EndDate() time.Time {
	return e.endDate
}

func (e EventDTO) NotificationBefore() time.Duration {
	return time.Duration(e.notificationBefore * 1e9)
}
