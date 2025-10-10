package logging_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilkerciblak/buldum-app/shared/logging"
)

var logger logging.ILogger

func TestLogger__WithField(t *testing.T) {

	testRequest := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader([]byte(`{"message":"test"}`)))
	tw := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {

		var request struct {
			Message string `json:"message"`
		}

		_ = json.NewDecoder(r.Body).Decode(&request)

		logger.SetLevel(logging.DEBUG)
		logger.WithField("message", request.Message)
		logger.LogSelf(request.Message)

	})

	mux.ServeHTTP(tw, testRequest)

}
