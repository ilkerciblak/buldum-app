package middleware_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilkerciblak/buldum-app/shared/middleware"
)

type TestRecorderType struct {
	ResponseWriter http.ResponseWriter
	Context        context.Context
}

func (t TestRecorderType) Header() http.Header {
	return t.ResponseWriter.Header()
}

func (t TestRecorderType) Write(b []byte) (int, error) {
	return t.ResponseWriter.Write(b)
}

func (t TestRecorderType) WriteHeader(statusCode int) {
	t.ResponseWriter.WriteHeader(statusCode)
}

func (t *TestRecorderType) WithContext(ctx context.Context) {
	t.Context = ctx
}

type Logging struct{}

type LogginMockMiddleware struct {
	Name string
}

func (m *LogginMockMiddleware) Act(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Mock LoggingMiddleware Before Serving HttpRequest")

		ctx := context.WithValue(r.Context(), &Logging{}, LogginMockMiddleware{Name: "mockito"})
		next.ServeHTTP(w, r.WithContext(ctx))

		log.Println("Mock LoggingMiddleware After Serving HttpRequest")
	}
}

type Author struct{}

type MockAuthMiddleware struct {
	Name string
}

type MockEndPoint struct {
}

func (m *MockAuthMiddleware) Act(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Mock AuthMiddleware Before Serving HttpRequest")

		ctx := context.WithValue(r.Context(), &Author{}, MockAuthMiddleware{Name: "mockito"})
		next.ServeHTTP(w, r.WithContext(ctx))

		log.Println("Mock AuthMiddleware After Serving HttpRequest")
	}
}

func TestSharedMiddleware__CreateMiddlewareChains(t *testing.T) {
	mux := http.NewServeMux()
	testRecorder := &TestRecorderType{
		ResponseWriter: httptest.NewRecorder(),
		Context:        context.Background(),
	}
	testRequest := httptest.NewRequest(http.MethodGet, "/test", nil)

	var someChain func(http.HandlerFunc) http.HandlerFunc = middleware.CreateMiddlewareChain(&MockAuthMiddleware{}, &LogginMockMiddleware{})

	next := someChain(func(w http.ResponseWriter, r *http.Request) {
		log.Println("EndPoint Served Its Duty")

		if a, b := w.(*TestRecorderType); b {
			a.WithContext(r.Context())
		}

	})

	mux.HandleFunc("/test", next)
	mux.ServeHTTP(testRecorder, testRequest)

	ctx := testRecorder.Context

	if _, exists := ctx.Value(&Author{}).(MockAuthMiddleware); !exists {
		t.Fatalf("Chaining Not Working As Expected")
	}

	if _, exists := ctx.Value(&Logging{}).(LogginMockMiddleware); !exists {
		t.Fatalf("Chaining Not Working As Expected Logging")
	}
}

func TestSharedMiddleware__ChainMiddlewareWithEndPoint(t *testing.T) {
	mux := http.NewServeMux()
	testRecorder := &TestRecorderType{
		ResponseWriter: httptest.NewRecorder(),
		Context:        context.Background(),
	}

	testRequest := httptest.NewRequest(http.MethodGet, "/test", nil)

	mux.HandleFunc(
		"/test",
		middleware.ChainMiddlewareWithEndPoint(
			func(w http.ResponseWriter, r *http.Request) {

				log.Println("EndPoint Served Its Duty")

				if a, b := w.(*TestRecorderType); b {
					a.WithContext(r.Context())
				}
			},
			&LogginMockMiddleware{},
			&MockAuthMiddleware{},
		))

	mux.ServeHTTP(testRecorder, testRequest)
	if _, exists := testRecorder.Context.Value(&Author{}).(MockAuthMiddleware); !exists {
		t.Fatalf("Chaining Not Working As Expected")
	}

	if _, exists := testRecorder.Context.Value(&Logging{}).(LogginMockMiddleware); !exists {
		t.Fatalf("Chaining Not Working As Expected Logging")

	}
}
