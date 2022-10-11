package test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/config"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/grpc/protobuf"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client protobuf.CalendarClient

func TestHandler(t *testing.T) {
	calendarConfig, err := config.NewConfig("../config/config_test.env")
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	conn, err := grpc.Dial(
		net.JoinHostPort(calendarConfig.GRPCHost, calendarConfig.GRPCPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("error connect to GRPC server:", err)
	}
	client = protobuf.NewCalendarClient(conn)

	startDate := time.Now().Add(3 * time.Second)
	startDateStr := startDate.Format(time.RFC3339)
	endDate := startDate.Add(1 * time.Second)
	endDateStr := endDate.Format(time.RFC3339)

	t.Run("Creat event", func(t *testing.T) {
		canNotAddNewEventForOtherUser(t, startDateStr, endDateStr)
		canAddNewEvent(t, startDateStr, endDateStr)
		canNotAddNewEventForSameTime(t, startDateStr, endDateStr)
		canAddNewEventForSameTimeForOtherUser(t, startDateStr, endDateStr)
	})

	t.Run("Update event", func(t *testing.T) {
		canNotEditEventForOtherUser(t, startDateStr, endDateStr)
		canEditEvent(t, startDateStr, endDateStr)
	})

	t.Run("Delete event", func(t *testing.T) {
		canNotDeleteEventForOtherUser(t)
		canDeleteEvent(t)
	})

	t.Run("Read event", func(t *testing.T) {
		canNotReadEventForOtherUser(t)
		canReadEvent(t)
		canReadEventForWeak(t, startDate)
	})

	t.Run("Send event", func(t *testing.T) {
		notificationForEventIsSent(t)
	})
}

func canNotAddNewEventForOtherUser(
	t *testing.T,
	startDate string,
	endDate string,
) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.CreateRequest{
		Title:              "Title",
		Description:        "Description",
		UserID:             1,
		StartDate:          startDate,
		EndDate:            endDate,
		NotificationBefore: 1,
		AuthorUserID:       2,
	}

	_, err := client.Create(ctx, r)
	require.Error(t, err)
	require.Equal(t, "rpc error: code = Unknown desc = read access is Denied", err.Error())
}

func canAddNewEvent(
	t *testing.T,
	startDate string,
	endDate string,
) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.CreateRequest{
		Title:              "Title",
		Description:        "Description",
		UserID:             1,
		StartDate:          startDate,
		EndDate:            endDate,
		NotificationBefore: 1,
		AuthorUserID:       1,
	}

	_, err := client.Create(ctx, r)
	require.NoError(t, err)
}

func canNotAddNewEventForSameTime(
	t *testing.T,
	startDate string,
	endDate string,
) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.CreateRequest{
		Title:              "Title 2",
		Description:        "Description 2",
		UserID:             1,
		StartDate:          startDate,
		EndDate:            endDate,
		NotificationBefore: 1,
		AuthorUserID:       1,
	}

	_, err := client.Create(ctx, r)
	require.Error(t, err)
	require.Equal(t, "rpc error: code = Unknown desc = the date is busy", err.Error())
}

func canAddNewEventForSameTimeForOtherUser(
	t *testing.T,
	startDate string,
	endDate string,
) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.CreateRequest{
		Title:              "Title",
		Description:        "Description",
		UserID:             2,
		StartDate:          startDate,
		EndDate:            endDate,
		NotificationBefore: 1,
		AuthorUserID:       2,
	}

	_, err := client.Create(ctx, r)
	require.NoError(t, err)
}

func canNotEditEventForOtherUser(t *testing.T, startDate string, endDate string) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.EditRequest{
		Id:                 1,
		Title:              "Title 2",
		Description:        "Description 2",
		StartDate:          startDate,
		EndDate:            endDate,
		NotificationBefore: 1,
		AuthorUserID:       2,
	}

	_, err := client.Edit(ctx, r)
	require.Error(t, err)
	require.Equal(t, "rpc error: code = Unknown desc = read access is Denied", err.Error())
}

func canEditEvent(t *testing.T, startDate string, endDate string) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.EditRequest{
		Id:                 1,
		Title:              "Title 2",
		Description:        "Description 2",
		StartDate:          startDate,
		EndDate:            endDate,
		NotificationBefore: 1,
		AuthorUserID:       1,
	}

	_, err := client.Edit(ctx, r)
	require.NoError(t, err)
}

func canNotDeleteEventForOtherUser(t *testing.T) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.DeleteRequest{
		ID:           1,
		AuthorUserID: 2,
	}

	_, err := client.Delete(ctx, r)
	require.Error(t, err)
	require.Equal(t, "rpc error: code = Unknown desc = read access is Denied", err.Error())
}

func canDeleteEvent(t *testing.T) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.DeleteRequest{
		ID:           2,
		AuthorUserID: 2,
	}

	_, err := client.Delete(ctx, r)
	require.NoError(t, err)
}

func canNotReadEventForOtherUser(t *testing.T) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.EventRequest{
		Id:           1,
		AuthorUserID: 2,
	}

	_, err := client.Event(ctx, r)
	require.Error(t, err)
	require.Equal(t, "rpc error: code = Unknown desc = read access is Denied", err.Error())
}

func canReadEvent(t *testing.T) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.EventRequest{
		Id:           1,
		AuthorUserID: 1,
	}

	response, err := client.Event(ctx, r)
	require.NoError(t, err)
	require.Equal(t, "Title 2", response.Title)
}

func notificationForEventIsSent(t *testing.T) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.EventRequest{
		Id:           1,
		AuthorUserID: 1,
	}

	time.Sleep(time.Second * 10) // waiting for the rabbit
	response, err := client.Event(ctx, r)
	require.NoError(t, err)
	require.Equal(t, true, response.IsSent)
}

func canReadEventForWeak(
	t *testing.T,
	date time.Time,
) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &protobuf.EventsRequest{
		Date:         date.Format("2006-01-02"),
		AuthorUserID: 1,
	}

	response, err := client.WeakEvents(ctx, r)
	require.NoError(t, err)
	require.Equal(t, 1, len(response.List))

	for _, eventResponse := range response.List {
		require.Equal(t, int32(1), eventResponse.Id)
		require.Equal(t, "Title 2", eventResponse.Title)
	}
}
