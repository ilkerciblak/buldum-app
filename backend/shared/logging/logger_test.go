package logging_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ilkerciblak/buldum-app/shared/logging"
)

var logger logging.ILogger

func TestLogger__TestSloggerImplementation(t *testing.T) {

	logger = logging.NewSlogger(logging.LoggerOptions{
		MinLevel:    logging.DEBUG,
		JsonLogging: true,
	})

	testRequest := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader([]byte(`{"message":"test"}`)))
	tw := httptest.NewRecorder()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		var startTime time.Time = time.Now()

		logger.With("start_time", startTime)
		time.Sleep(1 * time.Second)
		var request struct {
			Message string `json:"message"`
		}

		_ = json.NewDecoder(r.Body).Decode(&request)

		logger.With("request", request)
		duration := time.Since(startTime).Seconds()
		logger.WithGroup("group2", "1", "2")
		logger.Log(logging.INFO, r.Context(), "INFORMATION", "duration", duration)

	})

	mux.ServeHTTP(tw, testRequest)

}
