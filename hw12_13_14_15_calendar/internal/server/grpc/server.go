package internalgrpc

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/configs"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/app"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/domain/event"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/server/grpc/protobuf"
	"google.golang.org/grpc"
)

type Server struct {
	app        *app.App
	grpcServer *grpc.Server
	config     *configs.Config
	logger     logger.Logger
}

type CalendarServer struct {
	app    *app.App
	logger logger.Logger
	protobuf.UnimplementedCalendarServer
}

type EventRequest struct {
	Id                 int32
	Title              string
	Description        string
	StartDate          time.Time
	EndDate            time.Time
	NotificationBefore int32
}

type EventResponse struct {
	Id               int32
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

func NewServer(logger logger.Logger, app *app.App, config *configs.Config) *Server {
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
	lsn, err := net.Listen("tcp", net.JoinHostPort(s.config.GrpcHost, s.config.GrpcPort))
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

func (C CalendarServer) Create(ctx context.Context, request *protobuf.CreateRequest) (*protobuf.CreateResponse, error) {
	startTime, err := time.Parse(time.RFC3339, request.StartDate)
	if err != nil {
		C.logger.Error("failed parse startTime: " + err.Error())
		return nil, err
	}

	endTime, err := time.Parse(time.RFC3339, request.EndDate)
	if err != nil {
		C.logger.Error("failed parse endTime: " + err.Error())
		return nil, err
	}

	err = C.app.CreateEvent(
		request.Title,
		request.Description,
		int32(request.UserID),
		startTime,
		endTime,
		time.Duration(float64(request.NotificationBefore)*1e9),
		int32(request.AuthorUserID),
	)

	if err != nil {
		return nil, err
	}

	return &protobuf.CreateResponse{}, nil
}

func (C CalendarServer) Event(ctx context.Context, request *protobuf.EventRequest) (*protobuf.EventResponse, error) {
	eventEntity, err := C.app.ReadEvent(int32(request.Id), int32(request.AuthorUserID))
	if err != nil {
		return nil, err
	}

	t := convertEventToResponse(eventEntity)

	return t, nil
}

func (C CalendarServer) WeakEvents(ctx context.Context, request *protobuf.EventsRequest) (*protobuf.EventsResponse, error) {
	startTime, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		C.logger.Error("failed parse startTime: " + err.Error())
		return nil, err
	}

	events, err := C.app.FindEventsForWeek(startTime, 1)
	if err != nil {
		return nil, err
	}
	eventResponses := make(map[string]*protobuf.EventResponse)
	for i := range events {
		eventResponses[string(rune(i))] = convertEventToResponse(events[i])
	}

	return &protobuf.EventsResponse{List: eventResponses}, nil
}

func (C CalendarServer) Edit(ctx context.Context, request *protobuf.EditRequest) (*protobuf.EditResponse, error) {
	startTime, err := time.Parse(time.RFC3339, request.StartDate)
	if err != nil {
		C.logger.Error("failed parse startTime: " + err.Error())
		return nil, err
	}

	endTime, err := time.Parse(time.RFC3339, request.EndDate)
	if err != nil {
		C.logger.Error("failed parse endTime: " + err.Error())
		return nil, err
	}

	err = C.app.UpdateEvent(
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

func (C CalendarServer) Delete(ctx context.Context, request *protobuf.DeleteRequest) (*protobuf.DeleteResponse, error) {
	err := C.app.DeleteEvent(request.ID, request.AuthorUserID)
	if err != nil {
		return nil, err
	}

	return &protobuf.DeleteResponse{}, nil
}

func convertEventToResponse(eventEntity *event.Event) *protobuf.EventResponse {
	return &protobuf.EventResponse{
		Id:                 eventEntity.Id(),
		Title:              eventEntity.Title(),
		Description:        eventEntity.Description(),
		UserID:             eventEntity.UserID(),
		StartDate:          eventEntity.StartDate().Format(time.RFC3339),
		EndDate:            eventEntity.EndDate().Format(time.RFC3339),
		NotificationBefore: int32(eventEntity.NotificationBefore()),
	}
}
