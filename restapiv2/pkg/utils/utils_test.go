package utils

import (
	"net/http"
	"testing"
)

func TestMethodIsGetPutDelete(t *testing.T) {
	for _, method := range GetPutDeleteMethods {
		if MethodIsGetPutDelete(method) != true {
			t.Errorf("expected true, got false with method %s", method)
		}
	}
	if MethodIsGetPutDelete(http.MethodPost) != false {
		t.Error("expected true, got false with method POST")
	}
}
