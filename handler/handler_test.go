package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckinHandler(t *testing.T) {
	var (
		body       = bytes.NewBufferString(`{"name": "iporsut", "league": "dtac"}`)
		request, _ = http.NewRequest("POST", "/api/coach/checkin", body)
		recorder   = httptest.NewRecorder()
	)

	CheckinHandler(recorder, request)

	var (
		expectedJSON = `{"id":"53f87e7ad18a68e0a884d31e"}`
		actualJSON   = recorder.Body.String()
	)
	if recorder.Body.String() != expectedJSON {
		t.Errorf("expected %s but was %s", expectedJSON, actualJSON)
	}

	testStatusCode(t, recorder, http.StatusCreated)
	testContentTypeJSON(t, recorder)
}

func TestCheckoutHandler(t *testing.T) {
	var (
		body       = bytes.NewBufferString(`{"name": "iporsut", "league": "dtac"}`)
		request, _ = http.NewRequest("POST", "/api/coach/checkout", body)
		recorder   = httptest.NewRecorder()
	)

	CheckoutHandler(recorder, request)

	testStatusCode(t, recorder, http.StatusAccepted)
	testContentTypeJSON(t, recorder)
}

func testContentTypeJSON(t *testing.T, recorder *httptest.ResponseRecorder) {
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json but was %s", recorder.Header().Get("Content-Type"))
	}
}

func testStatusCode(t *testing.T, recorder *httptest.ResponseRecorder, code int) {
	if recorder.Code != code {
		t.Errorf("expect %d but was %d", code, recorder.Code)
	}
}
