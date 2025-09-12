package jsonmapper

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"

	"testing"
)

func TestJsonMapper_DecodeRequestBody(t *testing.T) {

	type testStruct struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		IsBald bool   `json:"isbald"`
	}

	cases := []struct {
		Name            string
		Input           *http.Request
		Expected        any
		DoesExpectError bool
	}{
		{
			Name:            "Decoding Empty Request Body Should Raise Error",
			DoesExpectError: true,
			Input:           httptest.NewRequest("POST", "/", nil),
			Expected:        nil,
		},
		{
			Name:            "Decoding Invalid Request Body Should Raise Error",
			DoesExpectError: true,
			Input:           httptest.NewRequest("POST", "/", bytes.NewReader([]byte(`{"invalid"}`))),
			Expected:        nil,
		},
		{
			Name:            "Decoding Whole Json Body Should Succeed",
			DoesExpectError: false,
			Input:           httptest.NewRequest("POST", "/", bytes.NewReader([]byte(`{"name":"ilkerciblak","age":30,"isbald":true}`))),
			Expected:        testStruct{Name: "ilkerciblak", Age: 30, IsBald: true},
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				output, err := DecodeRequestBody[testStruct](tc.Input)
				if tc.DoesExpectError {
					if err == nil {
						t.Fatalf("Test was expecting an error to occur")
					}
				} else {
					if err != nil {
						t.Fatalf("Test was not expecting an error while got: %v", err)
					}
				}

				if !reflect.DeepEqual(tc.Expected, output) && !tc.DoesExpectError {
					t.Fatalf("Expected output and result not comparable or deeply equal\nGot %v\tExpected %v", output, tc.Expected)
				}
			},
		)
	}

}

func TestJsonMapper_Encoder(t *testing.T) {

}
