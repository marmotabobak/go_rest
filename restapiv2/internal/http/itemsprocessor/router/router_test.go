package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"restapiv2/pkg/utils"
)

type TestCase struct {
	methods []string
	path string
	statusCode int
	result string
}

func TestNewItemsProcessorRouter(t *testing.T) {

	tests := []TestCase{
		// negative tests
		{
			methods: []string{http.MethodPost},
			path: "/stat",
			statusCode: http.StatusMethodNotAllowed,
			result: "action should be defined for POST method\n",
		},
		{
			methods: []string{http.MethodPost},
			path: "/item/1",
			statusCode: http.StatusMethodNotAllowed,
			result: "action should be defined for POST method\n",
		},
		{
			methods: utils.GetPutDeleteMethods[:],
			path: "/item/1/action",
			statusCode: http.StatusMethodNotAllowed,
			result: "GET/PUT/DELETE methods are not allowed for actions\n",
		},
		{
			methods: utils.PutDeleteMethods[:],
			path: "/stat",
			statusCode: http.StatusMethodNotAllowed,
			result: "stat action has only GET method\n",
		},
		// positive tests
		{
			methods: utils.GetPutDeleteMethods[:],
			path: "/item/1",
			statusCode: http.StatusOK,
			result: "",
		},
		{
			methods: []string{http.MethodGet},
			path: "/stat",
			statusCode: http.StatusOK,
			result: "",
		},
		{
			methods: []string{http.MethodPost},
			path: "/item/1/action",
			statusCode: http.StatusOK,
			result: "",
		},
	}

	router := NewItemsProcessorRouter()	

	for _, testCase := range tests {
		
		for _, method := range testCase.methods {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(method, testCase.path, nil)

			router.ServeHTTP(w, r)
			res := w.Result()
		
			if res.StatusCode != testCase.statusCode {
				t.Errorf("%d status code expected, got: %d", testCase.statusCode, res.StatusCode)
			}

			buf, _ := io.ReadAll(res.Body)
			defer res.Body.Close()
			if string(buf) != testCase.result {
				t.Errorf("for path %s and method %s \"%s\" result expected, got: \"%s\"", testCase.path, method, testCase.result, string(buf))
			}
		}
	
	// 	buf, _ := io.ReadAll(res.Body)
	// 	defer res.Body.Close()
	// }
		
	}


	// for _, method := range utils.AllHTTPMethods {


}