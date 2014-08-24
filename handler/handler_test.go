package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockTracker struct {
	t    *testing.T
	call int
}

func (m *mockTracker) CheckIn(user User) Checkin {
	if user.Name != "iporsut" {
		m.t.Errorf("expect name is iporsut")
	}

	if user.League != "dtac" {
		m.t.Errorf("expect league is league")
	}

	m.call++

	return Checkin{
		ID: "53f87e7ad18a68e0a884d31e",
	}
}

func (m *mockTracker) CheckOut(user User) error {
	if user.Name != "iporsut" {
		m.t.Errorf("expect name is iporsut")
	}

	if user.League != "dtac" {
		m.t.Errorf("expect league is league")
	}

	m.call++

	return nil
}

func (m *mockTracker) expectCallOnce() {
	if m.call != 1 {
		m.t.Errorf("expect Tracker must be call once time")
	}
}

func TestCheckinHandler(t *testing.T) {
	var (
		body       = bytes.NewBufferString(`{"name": "iporsut", "league": "dtac"}`)
		request, _ = http.NewRequest("POST", "/api/coach/checkin", body)
		recorder   = httptest.NewRecorder()
	)

	var m = &mockTracker{t: t}

	var trackerHandler = &TimeTrackerHandler{
		Tracker: m,
	}

	trackerHandler.checkin(recorder, request)

	var (
		expectedJSON = `{"id":"53f87e7ad18a68e0a884d31e"}`
		actualJSON   = recorder.Body.String()
	)
	if recorder.Body.String() != expectedJSON {
		t.Errorf("expected %s but was %s", expectedJSON, actualJSON)
	}

	testStatusCode(t, recorder, http.StatusCreated)
	testContentTypeJSON(t, recorder)

	m.expectCallOnce()
}

func TestCheckoutHandler(t *testing.T) {
	var (
		body           = bytes.NewBufferString(`{"name": "iporsut", "league": "dtac"}`)
		request, _     = http.NewRequest("POST", "/api/coach/checkout", body)
		recorder       = httptest.NewRecorder()
		m              = &mockTracker{t: t}
		trackerHandler = &TimeTrackerHandler{Tracker: m}
	)

	trackerHandler.checkout(recorder, request)

	testStatusCode(t, recorder, http.StatusAccepted)
	testContentTypeJSON(t, recorder)

	m.expectCallOnce()
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

func TestCheckinURLHandler(t *testing.T) {
	var (
		m          = &mockTracker{t: t}
		tracker    = TimeTrackerHandler{m}
		body       = bytes.NewBufferString(`{"name": "iporsut", "league": "dtac"}`)
		request, _ = http.NewRequest("POST", "/api/coach/checkin", body)
		recorder   = httptest.NewRecorder()
	)

	tracker.ServeHTTP(recorder, request)

	m.expectCallOnce()
}
