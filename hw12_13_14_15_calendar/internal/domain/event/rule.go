package event

import (
	"errors"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

var (
	ErrDateBusy            = errors.New("The date is busy")
	ErrWriteAccessDenied   = errors.New("Write access is Denied")
	ErrReadAccessDenied    = errors.New("Read access is Denied")
	ErrStartDateGreaterEnd = errors.New("The start date is greater than the end date")
)

type Rules struct {
	storage storage.EventStorage
}

func NewRules(storage storage.EventStorage) *Rules {
	return &Rules{storage: storage}
}

func (u *Rules) CheckDate(event *Event) error {
	if event.startDate.After(event.endDate) {
		return ErrStartDateGreaterEnd
	}

	eventDTOs, err := u.storage.FindListByPeriod(event.StartDate(), event.EndDate(), event.UserID())
	if err != nil {
		return err
	}

	for _, eventDTO := range eventDTOs {
		if event.id == eventDTO.ID() {
			continue
		}
		return ErrDateBusy
	}

	return nil
}

func (u *Rules) CheckCreateAccess(event *Event, userID int32) error {
	if userID != event.UserID() {
		return ErrReadAccessDenied
	}

	return nil
}

func (u *Rules) CheckReadAccess(event *Event, userID int32) error {
	if userID != event.UserID() {
		return ErrReadAccessDenied
	}

	return nil
}

func (u *Rules) CheckEditAccess(event *Event, userID int32) error {
	if userID != event.UserID() {
		return ErrWriteAccessDenied
	}

	return nil
}

func (u *Rules) CheckDeleteAccess(event *Event, userID int32) error {
	if userID != event.UserID() {
		return ErrWriteAccessDenied
	}

	return nil
}
