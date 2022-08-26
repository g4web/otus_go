package memorystorage

import (
	"testing"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	s := New()
	userID := 1
	startDate := time.Now()
	endDate := startDate.AddDate(0, 1, 0)
	duration, _ := time.ParseDuration("1h")

	t.Run("Create", func(t *testing.T) {
		helperCreateTest(t, s, userID, startDate, endDate, duration)
	})

	t.Run("Read", func(t *testing.T) {
		helperReadTest(t, s, userID, startDate, endDate, duration)
	})

	t.Run("Update", func(t *testing.T) {
		helperUpdateTest(t, s, userID, startDate, endDate, duration)
	})

	t.Run("Delete", func(t *testing.T) {
		helperDeleteTest(t, s, userID, startDate, endDate, duration)
	})
}

func helperCreateTest(
	t *testing.T, s *Storage,
	userID int, startDate time.Time, endDate time.Time, duration time.Duration,
) {
	t.Helper()
	err := s.Insert(storage.NewEventDTO(
		0,
		"title",
		"desc",
		int32(userID),
		startDate,
		endDate,
		duration,
	))
	require.NoError(t, err)

	eventDTO, err := s.FindOneByID(1)
	require.NoError(t, err)
	require.Equal(t, eventDTO, storage.NewEventDTO(
		1,
		"title",
		"desc",
		int32(userID),
		startDate,
		endDate,
		duration,
	), "The saved event is not equal to the founded")

	err = s.Delete(1)
	require.NoError(t, err)
}

func helperReadTest(
	t *testing.T,
	s *Storage, userID int, startDate time.Time, endDate time.Time, duration time.Duration,
) {
	t.Helper()
	err := s.Insert(storage.NewEventDTO(
		0,
		"title",
		"desc",
		int32(userID),
		startDate,
		endDate,
		duration,
	))
	require.NoError(t, err)

	eventDTO, err := s.FindOneByID(1)
	require.NoError(t, err)
	require.Equal(t, eventDTO, storage.NewEventDTO(
		1,
		"title",
		"desc",
		int32(userID),
		startDate,
		endDate,
		duration,
	), "The saved event is not equal to the founded")

	foundedEventDTOs, err := s.FindListByPeriod(startDate, endDate, int32(userID))
	require.NoError(t, err)
	require.Equal(t, 1, len(foundedEventDTOs), "Not founded events by period")

	foundedEventDTOs, err = s.FindListByPeriod(
		startDate.AddDate(0, -1, -1),
		startDate.AddDate(0, 0, 1),
		int32(userID),
	)
	require.NoError(t, err)
	require.Equal(t, 1, len(foundedEventDTOs), "Not founded events by period")

	foundedEventDTOs, err = s.FindListByPeriod(
		startDate.AddDate(0, 0, 27),
		startDate.AddDate(0, 0, 28),
		int32(userID),
	)
	require.NoError(t, err)
	require.Equal(t, 1, len(foundedEventDTOs), "Not founded events by period")

	foundedEventDTOs, err = s.FindListByPeriod(
		endDate,
		endDate.AddDate(0, 0, 1),
		int32(userID),
	)
	require.NoError(t, err)
	require.Equal(t, 0, len(foundedEventDTOs), "Founded not existing events by period")

	foundedEventDTOs, err = s.FindListByPeriod(
		startDate.AddDate(0, -1, 0),
		startDate.AddDate(0, 0, -1),
		int32(userID),
	)
	require.NoError(t, err)
	require.Equal(t, 0, len(foundedEventDTOs), "founded not existing events by period")

	d, _ := time.ParseDuration("1h0ms")
	foundedEventDTOs, err = s.FindNotificationByPeriod(
		startDate.Add(-d),
		startDate,
	)
	require.NoError(t, err)
	require.Equal(t, 0, len(foundedEventDTOs), "founded not existing notifications by period")

	d, _ = time.ParseDuration("1h1ms")
	foundedEventDTOs, err = s.FindNotificationByPeriod(
		startDate.Add(-d),
		startDate,
	)
	require.NoError(t, err)
	require.Equal(t, 1, len(foundedEventDTOs), "not founded notifications by period")

	err = s.Delete(1)
	require.NoError(t, err)
}

func helperUpdateTest(
	t *testing.T,
	s *Storage, userID int, startDate time.Time, endDate time.Time, duration time.Duration,
) {
	t.Helper()
	err := s.Insert(storage.NewEventDTO(
		0,
		"title",
		"desc",
		int32(userID),
		startDate,
		endDate,
		duration,
	))
	require.NoError(t, err)

	newVersionEvent := storage.NewEventDTO(
		1,
		"title 2",
		"desc 2",
		1,
		startDate,
		endDate,
		duration,
	)
	err = s.Update(1, newVersionEvent)
	require.NoError(t, err)
	eventDTO, err := s.FindOneByID(1)
	require.NoError(t, err)
	require.Equal(t, eventDTO, newVersionEvent, "The updated event is not equal to the founded")

	require.Equal(t, false, eventDTO.NotificationIsSent())
	err = s.MarkNotificationAsSent(eventDTO.ID())
	require.NoError(t, err)
	require.Equal(t, true, eventDTO.NotificationIsSent())

	err = s.Delete(1)
	require.NoError(t, err)
}

func helperDeleteTest(
	t *testing.T, s *Storage, userID int, startDate time.Time, endDate time.Time, duration time.Duration,
) {
	t.Helper()
	err := s.Insert(storage.NewEventDTO(
		0,
		"title",
		"desc",
		int32(userID),
		startDate,
		endDate,
		duration,
	))
	require.NoError(t, err)

	err = s.Delete(1)
	require.NoError(t, err)

	eventDTO, err := s.FindOneByID(1)
	require.Error(t, err)
	require.Equal(t, (*storage.EventDTO)(nil), eventDTO, "")

	err = s.Insert(storage.NewEventDTO(
		0,
		"title",
		"desc",
		int32(userID),
		startDate,
		endDate,
		duration,
	))
	require.NoError(t, err)

	rowsAffected, err := s.DeleteOld(startDate)
	require.NoError(t, err)
	require.Equal(t, int32(0), rowsAffected)

	d, _ := time.ParseDuration("1ns")

	rowsAffected, err = s.DeleteOld(startDate.Add(d))
	require.NoError(t, err)
	require.Equal(t, int32(1), rowsAffected)

	eventDTO, err = s.FindOneByID(1)
	require.Error(t, err)
	require.Equal(t, (*storage.EventDTO)(nil), eventDTO, "")
}
