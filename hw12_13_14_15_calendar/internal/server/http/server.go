package internalhttp

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/configs"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/app"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
)

type Server struct {
	httpServer *http.Server
	app        *app.App
	logger     logger.Logger
	config     *configs.Config
}

func NewServer(logger logger.Logger, app *app.App, config *configs.Config) *Server {
	httpServer := &http.Server{Addr: net.JoinHostPort(config.HttpHost, config.HttpPort)}
	requestStatistic := NewRequestStatistic(logger)
	http.Handle("/hello-world", requestStatistic.Middleware(http.HandlerFunc(helloWorld)))
	return &Server{httpServer: httpServer, app: app, logger: logger, config: config}
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

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Otus!")
}
