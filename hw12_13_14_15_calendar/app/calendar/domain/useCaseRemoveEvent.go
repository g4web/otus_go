package domain

import (
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type UseCaseRemoveEvent struct {
	storage storage.EventStorage
	rules   *Rules
}

func NewUseCaseRemoveEvent(storage storage.EventStorage) *UseCaseRemoveEvent {
	rules := NewRules(storage)

	return &UseCaseRemoveEvent{storage: storage, rules: rules}
}

func (u *UseCaseRemoveEvent) Remove(e *Event, userID int32) error {
	if err := u.rules.CheckDeleteAccess(e, userID); err != nil {
		return err
	}

	return u.storage.Delete(e.ID())
}
