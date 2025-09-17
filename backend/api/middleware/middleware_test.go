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
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/core/presentation"
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

type Panic struct{}

type MockEndpoint struct{}

func (m MockEndpoint) HandleRequest(w http.ResponseWriter, r *http.Request) (presentation.ApiResult[any], coredomain.IApplicationError) {
	type request struct {
		Message string `json:"message"`
	}
	var req request
	req, err := jsonmapper.DecodeRequestBody[request](r)
	if err != nil {
		rerr := &coredomain.InternalServerError
		rerr.Message = err.Error()
		return presentation.ApiResult[any]{}, rerr
	}

	if strings.Contains(req.Message, "error") {
		return presentation.ApiResult[any]{}, &coredomain.MethodNotAllowed
	}

	if strings.EqualFold(req.Message, "panic") {
		if a, b := w.(*TestTypeResponseRecorder); b {

			a.WithContext(context.WithValue(r.Context(), &Panic{}, true))
		}
		panic("Test Panic")
	}

	if a, b := w.(*TestTypeResponseRecorder); b {
		a.WithContext(r.Context())
	}

	return presentation.ApiResult[any]{Data: req}, nil
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

func TestApiMiddleware__CreatingMiddlewareChains(t *testing.T) {

	testReader := httptest.NewRecorder()
	testRequest := httptest.NewRequest("POST", "/mockito", bytes.NewReader([]byte(`{"message":"Yikess"}`)))

	testTypeRecorder := TestTypeResponseRecorder{
		ResponseWriter: testReader,
		Ctx:            ctx,
	}

	authedChain := middleware.CreateMiddlewareChain(&MockAuthMiddleware{})
	authedChain(MockEndpoint{}, &MockLoggingMiddleware{Name: "mockito"}).ServeHTTP(&testTypeRecorder, testRequest)

	if reflect.DeepEqual(testTypeRecorder.Ctx, context.WithValue(context.WithValue(ctx, &Logger{}, &MockLoggingMiddleware{Name: "mockito"}), &Author{}, &MockAuthMiddleware{})) {
		t.Fatalf("Expected Context is Not Satisfied")
	}
}

func TestApiMiddleware__PanicRecoverMiddleware(t *testing.T) {
	testReader := httptest.NewRecorder()
	testRequest := httptest.NewRequest("GET", "/mockito", bytes.NewReader([]byte(`{"message":"panic"}`)))
	testResponseRecorder := TestTypeResponseRecorder{
		ResponseWriter: testReader,
		Ctx:            ctx,
	}

	middleware.PanicRecoverMiddleware{}.Act(func(w http.ResponseWriter, r *http.Request) {
		MockEndpoint{}.HandleRequest(w, r)
	}).ServeHTTP(&testResponseRecorder, testRequest)

	if x, exists := testResponseRecorder.Ctx.Value(&Panic{}).(bool); exists {
		if !x {
			t.Fatalf("PanicRecoverMiddleware not works as expected")
		}
	} else {
		t.Fatalf("PanicRecoverMiddleware not works as expected, there is no value as `panic` in context")
	}

}
