package app

import (
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/domain/event"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	eventStorage storage.EventStorage
	logger       logger.Logger
}

func New(logger logger.Logger, eventStorage storage.EventStorage) *App {
	return &App{eventStorage: eventStorage, logger: logger}
}

func (a *App) CreateEvent(
	title string,
	description string,
	userID int32,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
	authorUserID int32,
) error {
	useCase := event.NewUseCaseCreateEvent(a.eventStorage)

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

func (a *App) ReadEvent(
	id int32,
	userID int32,
) (*event.Event, error) {
	useCase := event.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEvent(id, userID)
}

func (a *App) FindEventsForDay(
	startDate time.Time,
	userID int32,
) ([]*event.Event, error) {
	useCase := event.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEventsForDay(startDate, userID)
}

func (a *App) FindEventsForWeek(
	startDate time.Time,
	userID int32,
) ([]*event.Event, error) {
	useCase := event.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEventsForWeek(startDate, userID)
}

func (a *App) FindEventsForMonth(
	startDate time.Time,
	userID int32,
) ([]*event.Event, error) {
	useCase := event.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEventsForMonth(startDate, userID)
}

func (a *App) UpdateEvent(
	id int32,
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
	authorUserID int32,
) error {
	useCase := event.NewUseCaseEditEvent(a.eventStorage)

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

func (a *App) DeleteEvent(id int32, authorUserID int32) error {
	eventForDelete, err := a.ReadEvent(id, authorUserID)
	if err != nil {
		return err
	}

	useCase := event.NewUseCaseRemoveEvent(a.eventStorage)

	return useCase.Remove(eventForDelete, authorUserID)
}
