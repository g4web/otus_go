package messagequeue

import (
	"time"
)

type EventDTO struct {
	ID          int32
	Title       string
	Description string
	UserID      int32
	StartDate   time.Time
	EndDate     time.Time
}
