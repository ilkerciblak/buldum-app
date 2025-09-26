package corepresentation_test

import (
	"maps"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ilkerciblak/buldum-app/shared/core/application"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

func TestCorePresentation__PathValuesMapperTest(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{name}/zuppa/{age}", func(w http.ResponseWriter, r *http.Request) {

	})

	type targetType struct {
		Name string `path:"name"`
		Age  int    `path:"age"`
	}

	type invalidType struct {
		Name string
	}

	cases := []struct {
		Name             string
		Input            *http.Request
		targetType       interface{}
		ExpectedOutput   map[string]string
		DoesExpectsError bool
	}{
		{
			Name:             "Valid Target Should OK",
			targetType:       targetType{},
			Input:            httptest.NewRequest("GET", "/test/ilkerciblak/zuppa/30", nil),
			ExpectedOutput:   map[string]string{"name": "ilkerciblak", "age": "30"},
			DoesExpectsError: false,
		},
		{
			Name:             "In-Valid Target Should Give Empty Map",
			targetType:       invalidType{},
			Input:            httptest.NewRequest("GET", "/test/ilkerciblak/zuppa/30", nil),
			ExpectedOutput:   map[string]string{},
			DoesExpectsError: false,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, c.Input)
				output := corepresentation.PathValuesMapper(c.Input, c.targetType)

				if !maps.Equal(c.ExpectedOutput, output) {
					t.Fatalf("Output Expectations are NOT Satisfied\nExpected %v\nGot%v", c.ExpectedOutput, output)
				}
			},
		)
	}

}

func TestCorePresentation__QueryValuesMapperTest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {})

	cases := []struct {
		Name             string
		Input            *http.Request
		targetType       interface{}
		ExpectedOutput   map[string]interface{}
		DoesExpectsError bool
	}{
		{
			Name:             "No Query Parameter Should Return Empty Map",
			Input:            httptest.NewRequest("GET", "/test", nil),
			ExpectedOutput:   map[string]interface{}{},
			DoesExpectsError: false,
			targetType: struct {
				f1 string `query:"f1"`
			}{},
		},
		{
			Name:  "All Single Query Parameters Should OK",
			Input: httptest.NewRequest("GET", "/test?f1=test1&f2=test2", nil),
			ExpectedOutput: map[string]interface{}{
				"f1": "test1",
				"f2": "test2",
			},
			DoesExpectsError: false,
			targetType: struct {
				f1 string `query:"f1"`
				f2 string `query:"f2"`
			}{},
		},
		{
			Name:  "Some Single Some Multiple Query Parameters Should OK",
			Input: httptest.NewRequest("GET", "/test?f1=test1&f1=test2&f2=test3", nil),
			ExpectedOutput: map[string]interface{}{
				"f1": []string{"test1", "test2"},
				"f2": "test3",
			},
			DoesExpectsError: false,
			targetType: struct {
				f1 []string `query:"f1"`
				f2 string   `query:"f2"`
			}{},
		},
		{
			Name:  "Some Single Some Multiple Query Parameters, Not All Parameters Have Struct Tag Should OK",
			Input: httptest.NewRequest("GET", "/test?f1=test1&f1=test2&f2=test3&f3=test4", nil),
			ExpectedOutput: map[string]interface{}{
				"f1": []string{"test1", "test2"},
				"f2": "test3",
			},
			DoesExpectsError: false,
			targetType: struct {
				f1 []string `query:"f1"`
				f2 string   `query:"f2"`
				f3 string
			}{},
		},
		{
			Name:  "Request Have an Arbitrary Query Field, Should OK",
			Input: httptest.NewRequest("GET", "/test?f1=test1&f1=test2&f2=test3&foo=bar", nil),
			ExpectedOutput: map[string]interface{}{
				"f1": []string{"test1", "test2"},
				"f2": "test3",
			},
			DoesExpectsError: false,
			targetType: struct {
				f1 []string `query:"f1"`
				f2 string   `query:"f2"`
				f3 string   `query:"f3"`
			}{},
		},
		{
			Name:  "Request Have Fields With List Options Not Strings Also An Arbitrary Fields, Should OK",
			Input: httptest.NewRequest("GET", "/test?f1=true&f2=1&f2=2&f2=3&foo=bar", nil),
			ExpectedOutput: map[string]interface{}{
				"f1": true,
				"f2": []int{1, 2, 3},
			},
			DoesExpectsError: false,
			targetType: struct {
				f1 bool   `query:"f1"`
				f2 []int  `query:"f2"`
				f3 []bool `query:"f3"`
			}{},
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, c.Input)
				output := corepresentation.QueryParametersMapper(c.Input, c.targetType)
				if !maps.EqualFunc(c.ExpectedOutput, output, reflect.DeepEqual) {
					t.Fatalf("Output Expectations are NOT Satisfied\nExpected %v\nGot %v", c.ExpectedOutput, output)
				}

			},
		)
	}
}

func TestCorePresentation__QueryValuesWithCommonQueryParameters(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {})

	cases := []struct {
		Name             string
		Input            *http.Request
		targetType       interface{}
		ExpectedOutput   map[string]interface{}
		DoesExpectsError bool
	}{
		{
			Name:  "Common Query Params",
			Input: httptest.NewRequest("GET", "/test?page=3&order=desc", nil),
			ExpectedOutput: map[string]interface{}{
				"page":  3,
				"order": "desc",
			},
			DoesExpectsError: false,
			targetType:       application.CommonQueryParameters{},
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, c.Input)
				output := corepresentation.QueryParametersMapper(c.Input, c.targetType)
				if !maps.EqualFunc(c.ExpectedOutput, output, reflect.DeepEqual) {
					t.Fatalf("Output Expectations are NOT Satisfied\nExpected %v\nGot %v", c.ExpectedOutput, output)
				}

			},
		)
	}
}

func TestCorePresentation__QueryValuesAndPathValuesMapper(t *testing.T) {
	cases := []struct {
		Name             string
		Input            *http.Request
		targetType       interface{}
		ExpectedOutput   map[string]interface{}
		DoesExpectsError bool
	}{
		{
			Name:  "Single Path and  Single Query Parameter Should OK",
			Input: httptest.NewRequest("GET", "/test/testValue/accounts/testValue2?foo=bar", http.NoBody),
			targetType: struct {
				TestVal string `path:"test_val"`
				Foo     string `query:"foo"`
			}{},
			ExpectedOutput: map[string]interface{}{
				"test_val": "testValue",
				"foo":      "bar",
			},
			DoesExpectsError: false,
		},
		{
			Name:  "Multiple Path and  Multiple Query Parameter with Arrays Should OK",
			Input: httptest.NewRequest("GET", "/test/testValue/accounts/testValue2?foo=bar&foo=zar", http.NoBody),
			targetType: struct {
				TestVal  string   `path:"test_val"`
				TestVal2 string   `path:"test_val2"`
				Foo      []string `query:"foo"`
			}{},
			ExpectedOutput: map[string]interface{}{
				"test_val":  "testValue",
				"test_val2": "testValue2",
				"foo":       []string{"bar", "zar"},
			},
			DoesExpectsError: false,
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/test/{test_val}/accounts/{test_val2}", func(w http.ResponseWriter, r *http.Request) {})

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, c.Input)

				output := corepresentation.QueryAndPathParametersMapper(c.Input, c.targetType)
				if !maps.EqualFunc(c.ExpectedOutput, output, reflect.DeepEqual) {
					t.Fatalf("Output Not Satisfies Expectations\nExpects %v\nGot %v", c.ExpectedOutput, output)
				}

			},
		)
	}
}
