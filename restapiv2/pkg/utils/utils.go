package utils

import "net/http"

var AllHTTPMethods = [4]string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}

var GetPutDeleteMethods = [3]string{
	http.MethodGet,
	http.MethodPut,
	http.MethodDelete,
}

var PutDeleteMethods = [2]string{
	http.MethodPut,
	http.MethodDelete,
}

func MethodIsGetPutDelete(method string) bool {
	for _, httpMethod := range GetPutDeleteMethods {
		if method == httpMethod {
			return true
		}
	}
	return false
}

func MethodIsPutDelete(method string) bool {
	for _, httpMethod := range PutDeleteMethods {
		if method == httpMethod {
			return true
		}
	}
	return false
}
