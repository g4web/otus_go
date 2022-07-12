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
	eventsDict     map[int]*storage.EventDTO
	userEventsDict map[int][]int
}

func New() *Storage {
	return &Storage{eventsDict: make(map[int]*storage.EventDTO), userEventsDict: make(map[int][]int)}
}

func (s *Storage) Insert(e *storage.EventDTO) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	newId := len(s.eventsDict) + 1
	e.SetId(newId)
	s.eventsDict[e.ID()] = e
	s.userEventsDict[e.UserID()] = append(s.userEventsDict[e.UserID()], e.ID())

	return nil
}

func (s *Storage) Update(id int, e *storage.EventDTO) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.eventsDict[id]; ok {

		s.eventsDict[id] = e

		return true, nil
	}

	return false, ErrEventNotFound
}

func (s *Storage) Delete(id int) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if event, ok := s.eventsDict[id]; ok {

		indexForDelete := 0
		for index, eventID := range s.userEventsDict[event.UserID()] {
			if eventID == id {
				indexForDelete = index
				break
			}
		}

		s.userEventsDict[event.UserID()] = append(
			s.userEventsDict[event.UserID()][:indexForDelete],
			s.userEventsDict[event.UserID()][indexForDelete+1:]...,
		)

		delete(s.eventsDict, id)

		return true, nil
	}

	return false, ErrEventNotFound
}

func (s *Storage) FindOneById(id int) (*storage.EventDTO, error) {
	if eventDTO, ok := s.eventsDict[id]; ok {
		return eventDTO, nil
	}

	return nil, ErrEventNotFound
}

func (s *Storage) FindListByPeriod(startDate time.Time, endDate time.Time, userID int) ([]*storage.EventDTO, error) {
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
