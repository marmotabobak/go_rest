package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	// "restapiv2/pkg/utils"
)

type TestCase struct {
	method string
	path string
	statusCode int
	result string
}

func TestNewItemsProcessorRouter(t *testing.T) {

	// negative tests
	tests := []TestCase{
		{
			method: http.MethodPost,
			path: "/stat",
			statusCode: http.StatusMethodNotAllowed,
			result: "action should be defined for POST method",
		},
		{
			method: http.MethodPost,
			path: "/item/1",
			statusCode: http.StatusMethodNotAllowed,
			result: "action should be defined for POST method",
		},
		{
			method: http.MethodGet,
			path: "/item/1/action",
			statusCode: http.StatusMethodNotAllowed,
			result: "GET/PUT/DELETE methods are not allowed for actions",
		},
		{
			method: http.MethodPut,
			path: "/item/1/action",
			statusCode: http.StatusMethodNotAllowed,
			result: "GET/PUT/DELETE methods are not allowed for actions",
		},
		{
			method: http.MethodDelete,
			path: "/item/1/action",
			statusCode: http.StatusMethodNotAllowed,
			result: "GET/PUT/DELETE methods are not allowed for actions",
		},
		{
			method: http.MethodPut,
			path: "/stat",
			statusCode: http.StatusMethodNotAllowed,
			result: "stat action has only GET method",
		},
		{
			method: http.MethodDelete,
			path: "/stat",
			statusCode: http.StatusMethodNotAllowed,
			result: "stat action has only GET method",
		},
	}

	router := NewItemsProcessorRouter()	

	for _, testCase := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(testCase.method, testCase.path, nil)

		router.ServeHTTP(w, r)
		res := w.Result()
	
		if res.StatusCode != testCase.statusCode {
			t.Errorf("%d status code expected, got: %d", testCase.statusCode, res.StatusCode)
		}

		buf, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		if string(buf) != testCase.result + "\n" {
			t.Errorf("\"%s\" result expected, got: \"%s\"", testCase.result, string(buf))
		}
	
	// 	buf, _ := io.ReadAll(res.Body)
	// 	defer res.Body.Close()
	// }
		
	}


	// for _, method := range utils.AllHTTPMethods {


}