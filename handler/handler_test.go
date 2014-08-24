package handler

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"

	"net/http/httptest"
)

func TestHandler(t *testing.T) {

	handler := rest.ResourceHandler{
		DisableJsonIndent: true,
		ErrorLogger:       log.New(ioutil.Discard, "", 0),
	}

	handler.SetRoutes(
		&rest.Route{"POST", "/api/coach/checkin",
			func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteHeader(http.StatusCreated)
				data := map[string]string{"id": "53f87e7ad18a68e0a884d31e"}
				w.WriteJson(data)
			},
		},
		&rest.Route{"POST", "/api/coach/checkout",
			func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteHeader(http.StatusAccepted)
			},
		},
	)

	recorded := test.RunRequest(t, &handler, test.MakeSimpleRequest(
		"POST", "http://www.sprint3r.com/api/coach/checkin", &map[string]string{"name": "iporsut", "league": "dtac"}))
	recorded.CodeIs(201)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{"id":"53f87e7ad18a68e0a884d31e"}`)

	recorded = test.RunRequest(t, &handler, test.MakeSimpleRequest(
		"POST", "http://www.sprint3r.com/api/coach/checkout", &map[string]string{"name": "iporsut", "league": "dtac"}))
	recorded.CodeIs(202)
	recorded.ContentTypeIsJson()
}

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
