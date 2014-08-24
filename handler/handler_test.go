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

	if recorder.Code != http.StatusCreated {
		t.Errorf("expect %d but was %d", http.StatusCreated, recorder.Code)
	}

	var (
		expectedJSON = `{"id":"53f87e7ad18a68e0a884d31e"}`
		actualJSON   = recorder.Body.String()
	)
	if recorder.Body.String() != expectedJSON {
		t.Errorf("expected %s but was %s", expectedJSON, actualJSON)
	}
}

func TestCheckoutHandler(t *testing.T) {
	var (
		body       = bytes.NewBufferString(`{"name": "iporsut", "league": "dtac"}`)
		request, _ = http.NewRequest("POST", "/api/coach/checkout", body)
		recorder   = httptest.NewRecorder()
	)

	CheckoutHandler(recorder, request)

	if recorder.Code != http.StatusAccepted {
		t.Errorf("expect %d but was %d", http.StatusCreated, recorder.Code)
	}
}
