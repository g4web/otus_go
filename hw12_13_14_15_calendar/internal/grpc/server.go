package servergrpc

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/app/calendar"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/app/calendar/domain"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/config"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/grpc/protobuf"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"google.golang.org/grpc"
)

type Server struct {
	app        *calendar.Calendar
	grpcServer *grpc.Server
	config     *config.Config
	logger     logger.Logger
}

type CalendarServer struct {
	app    *calendar.Calendar
	logger logger.Logger
	protobuf.UnimplementedCalendarServer
}

type EventRequest struct {
	ID                 int32
	Title              string
	Description        string
	StartDate          time.Time
	EndDate            time.Time
	NotificationBefore int32
}

type EventResponse struct {
	ID               int32
	Title            string
	Description      string
	UserID           int32
	StartDate        time.Time
	EndDate          time.Time
	NotificationDate time.Time
}

type EventResponseList struct {
	Events []*EventResponse
}

func NewServer(logger logger.Logger, app *calendar.Calendar, config *config.Config) *Server {
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		RequestStatisticInterceptor(logger),
	))

	calendarServer := &CalendarServer{
		app:    app,
		logger: logger,
	}

	protobuf.RegisterCalendarServer(grpcServer, calendarServer)

	srv := &Server{
		app:        app,
		grpcServer: grpcServer,
		config:     config,
		logger:     logger,
	}

	return srv
}

func (s *Server) Start(ctx context.Context) error {
	lsn, err := net.Listen("tcp", net.JoinHostPort(s.config.GRPCHost, s.config.GRPCPort))
	if err != nil {
		s.logger.Error("fail start gprc server:" + err.Error())
	}

	if err := s.grpcServer.Serve(lsn); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("listen: " + err.Error())
		os.Exit(1)
	}
	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("gprc server is stopped")
	s.grpcServer.GracefulStop()

	return nil
}

func (c CalendarServer) Create(ctx context.Context, request *protobuf.CreateRequest) (*protobuf.CreateResponse, error) {
	startTime, err := time.Parse(time.RFC3339, request.StartDate)
	if err != nil {
		c.logger.Error("failed parse startTime: " + err.Error())
		return nil, err
	}

	endTime, err := time.Parse(time.RFC3339, request.EndDate)
	if err != nil {
		c.logger.Error("failed parse endTime: " + err.Error())
		return nil, err
	}

	err = c.app.CreateEvent(
		request.Title,
		request.Description,
		request.UserID,
		startTime,
		endTime,
		time.Duration(float64(request.NotificationBefore)*1e9),
		request.AuthorUserID,
	)

	if err != nil {
		return nil, err
	}

	return &protobuf.CreateResponse{}, nil
}

func (c CalendarServer) Event(ctx context.Context, request *protobuf.EventRequest) (*protobuf.EventResponse, error) {
	eventEntity, err := c.app.ReadEvent(request.Id, request.AuthorUserID)
	if err != nil {
		return nil, err
	}

	t := convertEventToResponse(eventEntity)

	return t, nil
}

func (c CalendarServer) WeakEvents(
	ctx context.Context,
	request *protobuf.EventsRequest,
) (*protobuf.EventsResponse, error) {
	startTime, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		c.logger.Error("failed parse startTime: " + err.Error())
		return nil, err
	}

	events, err := c.app.FindEventsForWeek(startTime, 1)
	if err != nil {
		return nil, err
	}
	eventResponses := make(map[string]*protobuf.EventResponse)
	for i := range events {
		eventResponses[string(rune(i))] = convertEventToResponse(events[i])
	}

	return &protobuf.EventsResponse{List: eventResponses}, nil
}

func (c CalendarServer) Edit(ctx context.Context, request *protobuf.EditRequest) (*protobuf.EditResponse, error) {
	startTime, err := time.Parse(time.RFC3339, request.StartDate)
	if err != nil {
		c.logger.Error("failed parse startTime: " + err.Error())
		return nil, err
	}

	endTime, err := time.Parse(time.RFC3339, request.EndDate)
	if err != nil {
		c.logger.Error("failed parse endTime: " + err.Error())
		return nil, err
	}

	err = c.app.UpdateEvent(
		request.Id,
		request.Title,
		request.Description,
		startTime,
		endTime,
		time.Duration(float64(request.NotificationBefore)*1e9),
		request.AuthorUserID,
	)

	if err != nil {
		return nil, err
	}

	return &protobuf.EditResponse{}, nil
}

func (c CalendarServer) Delete(ctx context.Context, request *protobuf.DeleteRequest) (*protobuf.DeleteResponse, error) {
	err := c.app.DeleteEvent(request.ID, request.AuthorUserID)
	if err != nil {
		return nil, err
	}

	return &protobuf.DeleteResponse{}, nil
}

func convertEventToResponse(eventEntity *domain.Event) *protobuf.EventResponse {
	return &protobuf.EventResponse{
		Id:                 eventEntity.ID(),
		Title:              eventEntity.Title(),
		Description:        eventEntity.Description(),
		UserID:             eventEntity.UserID(),
		StartDate:          eventEntity.StartDate().Format(time.RFC3339),
		EndDate:            eventEntity.EndDate().Format(time.RFC3339),
		NotificationBefore: int32(eventEntity.NotificationBefore()),
	}
}
