package serverhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/app/calendar"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/app/calendar/domain"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
)

type Handler struct {
	app    *calendar.Calendar
	logger logger.Logger
}

func NewHandler(app *calendar.Calendar, logger logger.Logger) *Handler {
	return &Handler{app: app, logger: logger}
}

func (h Handler) App() *calendar.Calendar {
	return h.app
}

func (h Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.SendSuccessResponse(w, HealthCheckMsg)
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var eventRequest EventRequest
	if err := decoder.Decode(&eventRequest); err != nil {
		h.SendBadResponse(w, http.StatusBadRequest, "failed decode eventRequest "+err.Error())
		return
	}

	err := h.app.CreateEvent(
		eventRequest.Title,
		eventRequest.Description,
		1,
		eventRequest.StartDate,
		eventRequest.EndDate,
		time.Duration(float64(eventRequest.NotificationBefore)*1e9),
		1,
	)
	if err != nil {
		h.SendBadResponse(w, http.StatusInternalServerError, "failed create eventRequest: "+err.Error())
		return
	}

	h.SendSuccessResponse(w, "")
}

func (h Handler) ReadOne(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.SendBadResponse(w, http.StatusBadRequest, "failed parse form: "+err.Error())
		return
	}

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 32)
	if err != nil {
		h.SendBadResponse(w, http.StatusBadRequest, "failed read id: "+err.Error())
		return
	}

	eventEntity, err := h.app.ReadEvent(int32(id), 1)

	var responseBytes []byte
	if err == nil {
		eventResponse := convertEventToResponse(eventEntity)
		responseBytes, err = json.Marshal(eventResponse)
	}

	if err != nil {
		h.SendBadResponse(w, http.StatusInternalServerError, "failed read eventEntity: "+err.Error())
	}

	h.SendSuccessResponse(w, string(responseBytes))
}

func (h Handler) ReadForWeak(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.SendBadResponse(w, http.StatusBadRequest, "failed parse form: "+err.Error())
		return
	}

	startTime, err := time.Parse("2006-01-02", r.Form.Get("start-date"))
	if err != nil {
		h.SendBadResponse(w, http.StatusBadRequest, "failed parse startTime: "+err.Error())
		return
	}

	events, err := h.app.FindEventsForWeek(startTime, 1)
	if err != nil {
		h.SendBadResponse(w, http.StatusNotFound, "failed find events: "+err.Error())
		return
	}

	eventResponses := make([]*EventResponse, len(events))
	for i := range events {
		eventResponses[i] = convertEventToResponse(events[i])
	}

	eventResponseList := &EventResponseList{eventResponses}

	responseBytes, err := json.Marshal(eventResponseList)
	if err != nil {
		h.SendBadResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.SendSuccessResponse(w, string(responseBytes))
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var eventRequest EventRequest
	if err := decoder.Decode(&eventRequest); err != nil {
		h.SendBadResponse(w, http.StatusBadRequest, "failed decode eventRequest: "+err.Error())
		return
	}
	err := h.app.UpdateEvent(
		eventRequest.ID,
		eventRequest.Title,
		eventRequest.Description,
		eventRequest.StartDate,
		eventRequest.EndDate,
		time.Duration(float64(eventRequest.NotificationBefore)*1e9),
		1,
	)
	if err != nil {
		h.SendBadResponse(w, http.StatusInternalServerError, "failed Update eventRequest: "+err.Error())
		return
	}
	h.SendSuccessResponse(w, "")
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error("failed parse form: " + err.Error())
	}

	id, _ := strconv.ParseInt(r.Form.Get("id"), 10, 32)
	err := h.app.DeleteEvent(int32(id), 1)
	if err != nil {
		h.SendBadResponse(w, http.StatusInternalServerError, "failed Delete event: "+err.Error())
		return
	}
	h.SendSuccessResponse(w, "")
}

func convertEventToResponse(event *domain.Event) *EventResponse {
	eventResponse := &EventResponse{
		event.ID(),
		event.Title(),
		event.Description(),
		event.UserID(),
		event.StartDate(),
		event.EndDate(),
		event.NotificationDate(),
	}
	return eventResponse
}

func (h Handler) SendSuccessResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, msg)
}

func (h Handler) SendBadResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, msg)
}
