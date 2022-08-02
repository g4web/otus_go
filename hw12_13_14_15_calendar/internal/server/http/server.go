package internalhttp

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/configs"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/app"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	"github.com/gorilla/mux"
)

const (
	HealthCheckMsg  = "I'm fine!"
	HealthCheckPath = "/health-check"
	EventPath       = "/event"
	WeakEventPath   = "/weak-events"
)

type Server struct {
	httpServer *http.Server
	logger     logger.Logger
	config     *configs.Config
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
	requestStatistic := NewRequestStatistic(logger)

	router := mux.NewRouter()
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.HttpHost, config.HttpPort),
		Handler: router,
	}
	server := &Server{httpServer: httpServer, logger: logger, config: config}

	handler := NewHandler(app, logger)
	router.Use(requestStatistic.Middleware)
	router.HandleFunc(HealthCheckPath, handler.HealthCheck).Methods("GET")
	router.HandleFunc(EventPath, handler.Create).Methods("POST")
	router.HandleFunc(EventPath, handler.ReadOne).Methods("GET")
	router.HandleFunc(WeakEventPath, handler.ReadForWeak).Methods("GET")
	router.HandleFunc(EventPath, handler.Update).Methods("PUT")
	router.HandleFunc(EventPath, handler.Delete).Methods("DELETE")

	return server
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.Error("ListenAndServe(): " + err.Error())

		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		s.logger.Warning("Shutdown(): " + err.Error())
	}

	return err
}
