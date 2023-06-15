package testing

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Testcase struct {
	Name         string
	Payload      interface{}
	Params       gin.Params
	ExpectedCode int
	ExpectedData interface{}
}

func TestRun(testcases []Testcase, function func(context *gin.Context), method string, withData bool, t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)

			requestJSON, _ := json.Marshal(tc.Payload)
			request := httptest.NewRequest(method, "/does-not-matter", bytes.NewBuffer(requestJSON))
			if tc.Payload == nil {
				request = httptest.NewRequest(method, "/does-not-matter", nil)
			}

			context.Params = tc.Params
			context.Request = request

			function(context)

			if recorder.Code != tc.ExpectedCode {
				t.Errorf("Expected status code %d but got %d", tc.ExpectedCode, recorder.Code)
			}

			if !withData || tc.ExpectedData == nil {
				return
			}

			var response interface{}
			if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
				t.Errorf("Error unmarshalling response body: %v", err)
			}

			expectedJSON, _ := json.Marshal(tc.ExpectedData)
			actualJSON, _ := json.Marshal(response)

			assertJSONResponse(t, expectedJSON, actualJSON)
		})
	}
}

func assertJSONResponse(t *testing.T, expectedJSON, actualJSON []byte) {
	var expected, actual interface{}
	if err := json.Unmarshal(expectedJSON, &expected); err != nil {
		t.Fatalf("Error unmarshalling expected JSON: %v", err)
	}
	if err := json.Unmarshal(actualJSON, &actual); err != nil {
		t.Fatalf("Error unmarshalling actual JSON: %v", err)
	}

	expectedMap, expectedIsMap := expected.(map[string]interface{})
	actualMap, actualIsMap := actual.(map[string]interface{})

	if expectedIsMap && actualIsMap {
		delete(expectedMap, "token")
		delete(actualMap, "token")

		if !reflect.DeepEqual(expectedMap, actualMap) {
			t.Errorf("Expected JSON response %s, but got %s", expectedJSON, actualJSON)
		}
	} else {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected JSON response %s, but got %s", expectedJSON, actualJSON)
		}
	}
}
