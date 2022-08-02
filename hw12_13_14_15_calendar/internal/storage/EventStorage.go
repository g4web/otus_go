package storage

import "time"

type EventStorage interface {
	Insert(e *EventDTO) error
	Update(id int32, e *EventDTO) error
	Delete(id int32) error
	FindOneById(id int32) (*EventDTO, error)
	FindListByPeriod(startDate time.Time, endDate time.Time, userID int32) ([]*EventDTO, error)
}
