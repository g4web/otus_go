package storage

import "time"

type EventStorage interface {
	Insert(e *EventDTO) error
	Update(id int32, e *EventDTO) error
	MarkNotificationAsSent(id int32) error
	Delete(id int32) error
	DeleteOld(endDate time.Time) (int32, error)
	FindOneByID(id int32) (*EventDTO, error)
	FindListByPeriod(startDate time.Time, endDate time.Time, userID int32) ([]*EventDTO, error)
	FindNotificationByPeriod(startDate time.Time, endDate time.Time) ([]*EventDTO, error)
}
