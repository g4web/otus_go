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
	Title string,
	description string,
	userID int,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
	authorUserID int,
) error {
	useCase := event.NewUseCaseCreateEvent(a.eventStorage)

	err := useCase.CreateEvent(
		Title,
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
	id int,
	userID int,
) (*event.Event, error) {
	useCase := event.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEvent(id, userID)
}

func (a *App) FindEventsForDay(
	startDate time.Time,
	userID int,
) ([]*event.Event, error) {
	useCase := event.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEventsForDay(startDate, userID)
}

func (a *App) FindEventsForWeek(
	startDate time.Time,
	userID int,
) ([]*event.Event, error) {
	useCase := event.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEventsForWeek(startDate, userID)
}

func (a *App) FindEventsForMonth(
	startDate time.Time,
	userID int,
) ([]*event.Event, error) {
	useCase := event.NewUseCaseFindEvent(a.eventStorage)

	return useCase.FindEventsForMonth(startDate, userID)
}

func (a *App) UpdateEvent(
	id int,
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	notificationBefore time.Duration,
	authorUserID int,
) (bool, error) {
	useCase := event.NewUseCaseEditEvent(a.eventStorage)

	eventForUpdate, err := a.ReadEvent(id, authorUserID)
	if err != nil {
		return false, err
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

func (a *App) DeleteEvent(id int, authorUserID int) (bool, error) {
	eventForDelete, err := a.ReadEvent(id, authorUserID)
	if err != nil {
		return false, err
	}

	useCase := event.NewUseCaseRemoveEvent(a.eventStorage)

	return useCase.Remove(eventForDelete, authorUserID)
}
