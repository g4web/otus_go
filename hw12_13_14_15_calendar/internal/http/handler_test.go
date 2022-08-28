package serverhttp

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/app/calendar"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/config"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	memorystorage "github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/stretchr/testify/require"
)

const requestGood = `
{
	"title" :"title string API 13",
	"description" :        "description string API 13",
	"userID" :             1,
    "startDate" :          "2021-11-30T18:00:00Z", 
	"endDate" :            "2021-12-30T18:00:00Z",
	"notificationBefore" : 3600
}
`

const requestWitErrorDates = `
{
	"title" :"title string API 13",
	"description" :        "description string API 13",
	"userID" :             1,
    "startDate" :          "2021-11-30T18:00:00Z", 
	"endDate" :            "2021-10-30T18:00:00Z",
	"notificationBefore" : 3600
}
`

const requestForUpdate = `
{
    "id": 1,
	"title" :"title string API 13",
	"description" :        "description string API 13",
    "startDate" :          "2021-11-30T18:00:00Z", 
	"endDate" :            "2021-12-30T18:00:00Z",
	"notificationBefore" : 3600
}
`

const expectedRequest = "{\"ID\":1,\"Title\":\"title string API 13\",\"Description\":\"description string API 13\"," +
	"\"UserID\":1,\"StartDate\":\"2021-11-30T18:00:00Z\",\"EndDate\":\"2021-12-30T18:00:00Z\",\"NotificationDate\"" +
	":\"2021-11-30T19:00:00Z\"}"

func TestHandler(t *testing.T) {
	calendarConfig, err := config.NewConfig("../../config/config_test.env")
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	logg := logger.New(calendarConfig.LogLevel, calendarConfig.LogFile)
	defer logg.Close()
	eventStorage := memorystorage.New()
	if err != nil {
		logg.Error(err.Error())
	}
	calendarApp := calendar.New(logg, eventStorage)
	handler := NewHandler(calendarApp, logg)
	hostPort := net.JoinHostPort(calendarConfig.HTTPHost, calendarConfig.HTTPPort)

	t.Run("HealthCheck", func(t *testing.T) {
		u := url.URL{Host: hostPort, Path: HealthCheckPath}
		r := httptest.NewRequest("GET", "http:"+u.String(), nil)
		w := httptest.NewRecorder()
		handler.HealthCheck(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		require.Equal(t, HealthCheckMsg, string(body))
	})

	t.Run("Create success", func(t *testing.T) {
		u := url.URL{Host: hostPort, Path: EventPath}
		r := httptest.NewRequest("POST", "http:"+u.String(), strings.NewReader(requestGood))
		w := httptest.NewRecorder()
		handler.Create(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Create error", func(t *testing.T) {
		u := url.URL{Host: hostPort, Path: EventPath}
		r := httptest.NewRequest("POST", "http:"+u.String(), strings.NewReader(requestWitErrorDates))
		w := httptest.NewRecorder()
		handler.Create(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		require.Equal(t, "failed create eventRequest: the start date is greater than the end date", string(body))
	})

	t.Run("Update success", func(t *testing.T) {
		u := url.URL{Host: hostPort, Path: EventPath}

		r := httptest.NewRequest("PUT", "http:"+u.String(), strings.NewReader(requestForUpdate))
		w := httptest.NewRecorder()
		handler.Update(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Update error", func(t *testing.T) {
		u := url.URL{Host: hostPort, Path: EventPath}
		r := httptest.NewRequest("PUT", "http:"+u.String(), strings.NewReader(requestWitErrorDates))
		w := httptest.NewRecorder()
		handler.Update(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		require.Equal(t, "failed Update eventRequest: event not found", string(body))
	})

	t.Run("Read success", func(t *testing.T) {
		u := url.URL{Host: hostPort, Path: EventPath, RawQuery: "id=1"}
		r := httptest.NewRequest("GET", "http:"+u.String(), strings.NewReader(""))
		w := httptest.NewRecorder()
		handler.ReadOne(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		require.Equal(t, expectedRequest, string(body))
	})

	t.Run("Read error", func(t *testing.T) {
		u := url.URL{Host: hostPort, Path: EventPath, RawQuery: "id=2"}
		r := httptest.NewRequest("GET", "http:"+u.String(), strings.NewReader(""))
		w := httptest.NewRecorder()
		handler.ReadOne(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		require.Equal(t, "failed read eventEntity: event not found", string(body))
	})

	t.Run("Delete error", func(t *testing.T) {
		u := url.URL{Host: hostPort, Path: EventPath, RawQuery: "id=2"}
		r := httptest.NewRequest("DELETE", "http:"+u.String(), strings.NewReader(""))
		w := httptest.NewRecorder()
		handler.Delete(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		require.Equal(t, "failed Delete event: event not found", string(body))
	})

	t.Run("Delete success", func(t *testing.T) {
		u := url.URL{Host: hostPort, Path: EventPath, RawQuery: "id=1"}
		r := httptest.NewRequest("DELETE", "http:"+u.String(), strings.NewReader(""))
		w := httptest.NewRecorder()
		handler.Delete(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
