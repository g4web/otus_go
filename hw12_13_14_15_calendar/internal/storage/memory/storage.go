package memorystorage

import (
	"errors"
	"sync"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
)

var ErrEventNotFound = errors.New("event not found")

type Storage struct {
	mu             sync.RWMutex
	eventsDict     map[int32]*storage.EventDTO
	userEventsDict map[int32][]int32
}

func New() *Storage {
	return &Storage{eventsDict: make(map[int32]*storage.EventDTO), userEventsDict: make(map[int32][]int32)}
}

func (s *Storage) Insert(e *storage.EventDTO) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	newID := len(s.eventsDict) + 1
	e.SetID(int32(newID))
	s.eventsDict[e.ID()] = e
	s.userEventsDict[e.UserID()] = append(s.userEventsDict[e.UserID()], e.ID())

	return nil
}

func (s *Storage) Update(id int32, e *storage.EventDTO) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.eventsDict[id]; ok {
		s.eventsDict[id] = e

		return nil
	}

	return ErrEventNotFound
}

func (s *Storage) MarkNotificationAsSent(id int32) error {
	if eventDTO, ok := s.eventsDict[id]; ok {
		eventDTO.MarkNotificationAsSent()
		return nil
	}

	return ErrEventNotFound
}

func (s *Storage) Delete(id int32) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if eventDTO, ok := s.eventsDict[id]; ok {
		s.deleteEvent(eventDTO)
		return nil
	}

	return ErrEventNotFound
}

func (s *Storage) DeleteOld(endDate time.Time) (int32, error) {
	rowsAffected := 0
	for _, eventDTO := range s.eventsDict {
		if eventDTO.StartDate().Before(endDate) {
			s.deleteEvent(eventDTO)
			rowsAffected++
		}
	}

	return int32(rowsAffected), nil
}

func (s *Storage) FindOneByID(id int32) (*storage.EventDTO, error) {
	if eventDTO, ok := s.eventsDict[id]; ok {
		return eventDTO, nil
	}

	return nil, ErrEventNotFound
}

func (s *Storage) FindListByPeriod(startDate time.Time, endDate time.Time, userID int32) ([]*storage.EventDTO, error) {
	var result []*storage.EventDTO
	if userEventIds, ok := s.userEventsDict[userID]; ok {
		for _, eventID := range userEventIds {
			eventDTO, ok := s.eventsDict[eventID]
			if !ok {
				continue
			}
			if eventDTO.StartDate().Before(endDate) && eventDTO.EndDate().After(startDate) {
				result = append(result, eventDTO)
			}
		}
	}

	return result, nil
}

func (s *Storage) FindNotificationByPeriod(startDate time.Time, endDate time.Time) ([]*storage.EventDTO, error) {
	var result []*storage.EventDTO

	for _, eventDTO := range s.eventsDict {
		notificationTime := eventDTO.StartDate().Add(-eventDTO.NotificationBefore())
		if !eventDTO.NotificationIsSent() && notificationTime.After(startDate) && notificationTime.Before(endDate) {
			result = append(result, eventDTO)
		}
	}

	return result, nil
}

func (s *Storage) deleteEvent(eventDTO *storage.EventDTO) {
	indexForDelete := 0
	for index, eventID := range s.userEventsDict[eventDTO.UserID()] {
		if eventID == eventDTO.ID() {
			indexForDelete = index
			break
		}
	}

	s.userEventsDict[eventDTO.UserID()] = append(
		s.userEventsDict[eventDTO.UserID()][:indexForDelete],
		s.userEventsDict[eventDTO.UserID()][indexForDelete+1:]...,
	)

	delete(s.eventsDict, eventDTO.ID())
}
