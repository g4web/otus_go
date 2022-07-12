package storage

import (
	"time"
)

type EventStorage interface {
	Insert(e *EventDTO) error
	Update(id int, e *EventDTO) (bool, error)
	Delete(id int) (bool, error)
	FindOneById(id int) (*EventDTO, error)
	FindListByPeriod(startDate time.Time, endDate time.Time, userID int) ([]*EventDTO, error)
}
