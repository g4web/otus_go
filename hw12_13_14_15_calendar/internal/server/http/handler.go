package internalhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/app"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/domain/event"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
)

type Handler struct {
	app    *app.App
	logger logger.Logger
}

func NewHandler(app *app.App, logger logger.Logger) *Handler {
	return &Handler{app: app, logger: logger}
}

func (h Handler) App() *app.App {
	return h.app
}

func (s Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	s.SendSuccessResponse(w, HealthCheckMsg)
}

func (s Handler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var eventRequest EventRequest
	if err := decoder.Decode(&eventRequest); err != nil {
		s.SendBadResponse(w, http.StatusBadRequest, "failed decode eventRequest "+err.Error())
		return
	}

	err := s.app.CreateEvent(
		eventRequest.Title,
		eventRequest.Description,
		1,
		eventRequest.StartDate,
		eventRequest.EndDate,
		time.Duration(float64(eventRequest.NotificationBefore)*1e9),
		1,
	)
	if err != nil {
		s.SendBadResponse(w, http.StatusInternalServerError, "failed create eventRequest: "+err.Error())
		return
	}

	s.SendSuccessResponse(w, "")
}

func (s Handler) ReadOne(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.SendBadResponse(w, http.StatusBadRequest, "failed parse form: "+err.Error())
		return
	}

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 32)
	if err != nil {
		s.SendBadResponse(w, http.StatusBadRequest, "failed read id: "+err.Error())
		return
	}

	eventEntity, err := s.app.ReadEvent(int32(id), 1)

	var responseBytes []byte
	if err == nil {
		eventResponse := convertEventToResponse(eventEntity)
		responseBytes, err = json.Marshal(eventResponse)
	}

	if err != nil {
		s.SendBadResponse(w, http.StatusInternalServerError, "failed read eventEntity: "+err.Error())
	}

	s.SendSuccessResponse(w, string(responseBytes))
}

func (s Handler) ReadForWeak(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.SendBadResponse(w, http.StatusBadRequest, "failed parse form: "+err.Error())
		return
	}

	startTime, err := time.Parse("2006-01-02", r.Form.Get("start-date"))
	if err != nil {
		s.SendBadResponse(w, http.StatusBadRequest, "failed parse startTime: "+err.Error())
		return
	}

	events, err := s.app.FindEventsForWeek(startTime, 1)
	if err != nil {
		s.SendBadResponse(w, http.StatusNotFound, "failed find events: "+err.Error())
		return
	}

	eventResponses := make([]*EventResponse, len(events))
	for i := range events {
		eventResponses[i] = convertEventToResponse(events[i])
	}

	eventResponseList := &EventResponseList{eventResponses}

	responseBytes, err := json.Marshal(eventResponseList)
	if err != nil {
		s.SendBadResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	s.SendSuccessResponse(w, string(responseBytes))
}

func (s Handler) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var eventRequest EventRequest
	if err := decoder.Decode(&eventRequest); err != nil {
		s.SendBadResponse(w, http.StatusBadRequest, "failed decode eventRequest: "+err.Error())
		return
	}
	err := s.app.UpdateEvent(
		eventRequest.Id,
		eventRequest.Title,
		eventRequest.Description,
		eventRequest.StartDate,
		eventRequest.EndDate,
		time.Duration(float64(eventRequest.NotificationBefore)*1e9),
		1,
	)
	if err != nil {
		s.SendBadResponse(w, http.StatusInternalServerError, "failed Update eventRequest: "+err.Error())
		return
	}
	s.SendSuccessResponse(w, "")
}

func (s Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.logger.Error("failed parse form: " + err.Error())
	}

	id, _ := strconv.ParseInt(r.Form.Get("id"), 10, 32)
	err := s.app.DeleteEvent(int32(id), 1)
	if err != nil {
		s.SendBadResponse(w, http.StatusInternalServerError, "failed Delete event: "+err.Error())
		return
	}
	s.SendSuccessResponse(w, "")
}

func convertEventToResponse(event *event.Event) *EventResponse {
	eventResponse := &EventResponse{
		event.Id(),
		event.Title(),
		event.Description(),
		event.UserID(),
		event.StartDate(),
		event.EndDate(),
		event.NotificationDate(),
	}
	return eventResponse
}

func (s Handler) SendSuccessResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, msg)
}

func (s Handler) SendBadResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, msg)
}
