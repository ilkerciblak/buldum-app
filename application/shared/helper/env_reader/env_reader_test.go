package envreader

import (
	"os"
	"strings"
	"testing"
)

func TestEnvReader_GetStringOrDefault(t *testing.T) {
	portKey := "PORT"
	portVal := "3000"
	defaultValue := "DefaultValue"
	os.Setenv(portKey, portVal)

	cases := []struct {
		Name            string
		Input           string
		ExpectedOutput  string
		DoesExpectError bool
	}{
		{
			Name:            "Expecting Result to be Default Result",
			Input:           "NotExistingKey",
			ExpectedOutput:  defaultValue,
			DoesExpectError: false,
		},
		{
			Name:            "Expecting Result to be 3000",
			Input:           portKey,
			ExpectedOutput:  portVal,
			DoesExpectError: false,
		},
	}

	for _, testCase := range cases {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				output := GetStringOrDefault(testCase.Input, defaultValue)

				if strings.Compare(testCase.ExpectedOutput, output) != 0 {
					t.Errorf("Expected Output %s\tWhile Got: %s", testCase.ExpectedOutput, output)
				}
			},
		)
	}
}
