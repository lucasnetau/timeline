package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jeffbmartinez/timeline/handler"
)

func TestSimpleEndpoint_NoParams(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://example.com", nil)
	response := httptest.NewRecorder()

	handler.Simple(response, request)

	const EXPECTED_STATUS_CODE = http.StatusBadRequest

	if response.Code != EXPECTED_STATUS_CODE {
		t.Fatalf("Should have received status code %v", EXPECTED_STATUS_CODE)
	}
}

func TestSimpleEndpoint_MissingSeriesParam(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost?key1=value1&key2=value", nil)
	response := httptest.NewRecorder()

	handler.Simple(response, request)

	const EXPECTED_STATUS_CODE = http.StatusBadRequest

	if response.Code != EXPECTED_STATUS_CODE {
		t.Fatalf("Should have received status code %v", EXPECTED_STATUS_CODE)
	}
}
