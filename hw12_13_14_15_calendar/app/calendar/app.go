package calendar

import (
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/app/calendar/domain"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type Calendar struct {
	eventStorage storage.EventStorage
	logger       logger.Logger
}

func New(logger logger.Logger, eventStorage storage.EventStorage) *Calendar {
	return &Calendar{eventStorage: eventStorage, logger: logger}
}

func (a *Calendar) CreateEvent(
	title string,
	description string,
	userID int32,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
	authorUserID int32,
) error {
	useCase := domain.NewUseCaseCreateEvent(a.eventStorage)

	err := useCase.CreateEvent(
		title,
		description,
		userID,
		startDate,
		endDate,
		notificationBefore,
		authorUserID,
	)

	return err
}

func (a *Calendar) ReadEvent(
	id int32,
	userID int32,
) (*domain.Event, error) {
	useCase := domain.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEvent(id, userID)
}

func (a *Calendar) FindEventsForDay(
	startDate time.Time,
	userID int32,
) ([]*domain.Event, error) {
	useCase := domain.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEventsForDay(startDate, userID)
}

func (a *Calendar) FindEventsForWeek(
	startDate time.Time,
	userID int32,
) ([]*domain.Event, error) {
	useCase := domain.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEventsForWeek(startDate, userID)
}

func (a *Calendar) FindEventsForMonth(
	startDate time.Time,
	userID int32,
) ([]*domain.Event, error) {
	useCase := domain.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEventsForMonth(startDate, userID)
}

func (a *Calendar) UpdateEvent(
	id int32,
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
	authorUserID int32,
) error {
	useCase := domain.NewUseCaseEditEvent(a.eventStorage)

	eventForUpdate, err := a.ReadEvent(id, authorUserID)
	if err != nil {
		return err
	}

	return useCase.EditEvent(
		eventForUpdate,
		title,
		description,
		startDate,
		endDate,
		notificationBefore,
		authorUserID,
	)
}

func (a *Calendar) DeleteEvent(id int32, authorUserID int32) error {
	eventForDelete, err := a.ReadEvent(id, authorUserID)
	if err != nil {
		return err
	}

	useCase := domain.NewUseCaseRemoveEvent(a.eventStorage)

	return useCase.Remove(eventForDelete, authorUserID)
}
