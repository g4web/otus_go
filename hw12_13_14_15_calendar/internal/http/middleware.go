package serverhttp

import (
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
)

type RequestStatistic struct {
	logger logger.Logger
}

func NewRequestStatistic(logger logger.Logger) *RequestStatistic {
	return &RequestStatistic{logger: logger}
}

func (rs *RequestStatistic) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := GetIP(r)
		start := time.Now()
		startStr := start.GoString()
		method := r.Method
		path := r.URL.Path
		query := r.RequestURI
		httpVersion := r.Proto
		userAgent := r.UserAgent()

		next.ServeHTTP(w, r)

		latency := time.Since(start).String()
		code := rs.getStatusCode(w)

		rs.logger.Info(
			ip +
				" " + startStr +
				" " + method +
				" " + path +
				" " + query +
				" " + httpVersion +
				" " + code +
				" " + latency +
				" " + userAgent)
	})
}

func (rs *RequestStatistic) getStatusCode(w http.ResponseWriter) string {
	codeInt := reflect.ValueOf(w).Elem().FieldByName("status").Int()
	code := strconv.Itoa(int(codeInt))
	return code
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
