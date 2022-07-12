package memorystorage

import (
	"testing"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	s := New()
	userId := 1
	startDate := time.Now()
	endDate := startDate.AddDate(0, 1, 0)
	duration, _ := time.ParseDuration("1h")

	t.Run("Create", func(t *testing.T) {
		err := s.Insert(storage.NewEventDTO(
			0,
			"title",
			"desc",
			userId,
			startDate,
			endDate,
			duration,
		))
		require.NoError(t, err)

		eventDTO, err := s.FindOneById(1)
		require.NoError(t, err)
		require.Equal(t, eventDTO, storage.NewEventDTO(
			1,
			"title",
			"desc",
			userId,
			startDate,
			endDate,
			duration,
		), "The saved event is not equal to the founded")

		result, err := s.Delete(1)
		require.Equal(t, true, result)
		require.NoError(t, err)
	})

	t.Run("Read", func(t *testing.T) {
		err := s.Insert(storage.NewEventDTO(
			0,
			"title",
			"desc",
			userId,
			startDate,
			endDate,
			duration,
		))
		require.NoError(t, err)

		eventDTO, err := s.FindOneById(1)
		require.NoError(t, err)
		require.Equal(t, eventDTO, storage.NewEventDTO(
			1,
			"title",
			"desc",
			userId,
			startDate,
			endDate,
			duration,
		), "The saved event is not equal to the founded")

		foundedEventDTOs, err := s.FindListByPeriod(startDate, endDate, userId)
		require.NoError(t, err)
		require.Equal(t, 1, len(foundedEventDTOs), "Not founded events by period")

		foundedEventDTOs, err = s.FindListByPeriod(
			startDate.AddDate(0, -1, -1),
			startDate.AddDate(0, 0, 1),
			userId,
		)
		require.NoError(t, err)
		require.Equal(t, 1, len(foundedEventDTOs), "Not founded events by period")

		foundedEventDTOs, err = s.FindListByPeriod(
			startDate.AddDate(0, 0, 27),
			startDate.AddDate(0, 0, 28),
			userId,
		)
		require.NoError(t, err)
		require.Equal(t, 1, len(foundedEventDTOs), "Not founded events by period")

		foundedEventDTOs, err = s.FindListByPeriod(
			endDate,
			endDate.AddDate(0, 0, 1),
			userId,
		)
		require.NoError(t, err)
		require.Equal(t, 0, len(foundedEventDTOs), "Founded not existing events by period")

		foundedEventDTOs, err = s.FindListByPeriod(
			startDate.AddDate(0, -1, 0),
			startDate.AddDate(0, 0, -1),
			userId,
		)
		require.NoError(t, err)
		require.Equal(t, 0, len(foundedEventDTOs), "Founded not existing events by period")

		result, err := s.Delete(1)
		require.Equal(t, true, result)
		require.NoError(t, err)
	})

	t.Run("Update", func(t *testing.T) {
		err := s.Insert(storage.NewEventDTO(
			0,
			"title",
			"desc",
			userId,
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
		result, err := s.Update(1, newVersionEvent)
		require.Equal(t, result, true, "")
		require.NoError(t, err)
		eventDTO, err := s.FindOneById(1)
		require.Equal(t, eventDTO, newVersionEvent, "The updated event is not equal to the founded")

		result, err = s.Delete(1)
		require.Equal(t, result, true, "")
		require.NoError(t, err)
	})

	t.Run("Delete", func(t *testing.T) {
		err := s.Insert(storage.NewEventDTO(
			0,
			"title",
			"desc",
			userId,
			startDate,
			endDate,
			duration,
		))
		require.NoError(t, err)

		result, err := s.Delete(1)
		require.Equal(t, true, result)
		require.NoError(t, err)

		eventDTO, err := s.FindOneById(1)
		require.Error(t, err)
		require.Equal(t, (*storage.EventDTO)(nil), eventDTO, "")
	})
}
