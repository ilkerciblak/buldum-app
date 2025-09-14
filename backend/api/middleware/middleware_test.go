package middleware_test

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/ilkerciblak/buldum-app/api/middleware"
	"github.com/ilkerciblak/buldum-app/shared/core/domain"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
)

type TestTypeResponseRecorder struct {
	ResponseWriter http.ResponseWriter
	Ctx            context.Context
}

func (t TestTypeResponseRecorder) Header() http.Header {
	return t.ResponseWriter.Header()
}

func (t TestTypeResponseRecorder) WriteHeader(statusCode int) {
	t.ResponseWriter.WriteHeader(statusCode)
}

func (t TestTypeResponseRecorder) Write(data []byte) (int, error) {
	return t.ResponseWriter.Write(data)
}

func (t *TestTypeResponseRecorder) WithContext(ctx context.Context) {
	t.Ctx = ctx
}

type Logger struct{}

type MockLoggingMiddleware struct {
	Name string
}

func (m *MockLoggingMiddleware) Act(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logging Before Running Next\n")

		ctx := context.WithValue(r.Context(), &Logger{}, MockLoggingMiddleware{Name: "mockito"})
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Printf("Logging After Running Next\n")
	}

}

type Author struct{}

type MockAuthMiddleware struct {
}

func (m *MockAuthMiddleware) Act(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Auth Middleware Before Running Next\n")

		ctx := context.WithValue(r.Context(), &Author{}, MockAuthMiddleware{})
		next.ServeHTTP(w, r.WithContext(ctx))

		log.Printf("Auth Middleware After Running Next\n")
	}
}

type MockEndpoint struct{}

func (m MockEndpoint) HandleRequest(w http.ResponseWriter, r *http.Request) (any, domain.IApplicationError) {
	type request struct {
		Message string `json:"message"`
	}
	var req request
	req, err := jsonmapper.DecodeRequestBody[request](r)
	if err != nil {
		rerr := &domain.InternalServerError
		rerr.Message = err.Error()
		return nil, rerr
	}

	if strings.Contains(req.Message, "error") {
		return nil, &domain.MethodNotAllowed
	}

	if a, b := w.(*TestTypeResponseRecorder); b {
		a.WithContext(r.Context())
	}

	return req, nil
}

func (m MockEndpoint) Path() string {
	return "/mockito"
}

var ctx context.Context = context.Background()

func TestApiMiddleware__ChainMiddleware(t *testing.T) {
	cases := []struct {
		Name            string
		DoesExpectError bool
		Input           []middleware.IMiddleware
		ExpectedOutput  []context.Context
	}{
		{
			Name: "Ananza xd",
			Input: []middleware.IMiddleware{&MockLoggingMiddleware{
				Name: "mockito",
			}, &MockAuthMiddleware{}},
			ExpectedOutput: []context.Context{
				context.WithValue(context.WithValue(ctx, &Logger{}, &MockLoggingMiddleware{Name: "mockito"}), &Author{}, &MockAuthMiddleware{}),
			},
		},
	}

	for i, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				testReader := httptest.NewRecorder()
				testRequest := httptest.NewRequest("POST", "/mockito", bytes.NewReader([]byte(`{"message":"Yikess"}`)))

				testTypeRecorder := TestTypeResponseRecorder{
					ResponseWriter: testReader,
					Ctx:            ctx,
				}
				middleware.ChainMiddlewaresWithEndpoint(MockEndpoint{}, tc.Input...).ServeHTTP(&testTypeRecorder, testRequest)

				if reflect.DeepEqual(testTypeRecorder.Ctx, tc.ExpectedOutput[i]) {
					t.Fatalf("Expected Context is Not Satisfied")
				}

			},
		)
	}

}
