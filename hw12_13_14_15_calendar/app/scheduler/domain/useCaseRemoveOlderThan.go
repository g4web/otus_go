package domain

import (
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type UseCaseRemoveOldEvents struct {
	storage storage.EventStorage
}

func NewUseCaseRemoveOldEvents(storage storage.EventStorage) *UseCaseRemoveOldEvents {
	return &UseCaseRemoveOldEvents{storage: storage}
}

func (u *UseCaseRemoveOldEvents) RemoveOlderThan(endDate time.Time) (int32, error) {
	return u.storage.DeleteOld(endDate)
}
