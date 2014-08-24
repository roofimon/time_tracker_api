package handler_test

import (
	"io/ioutil"
	"log"
	"testing"
	h "time_tracker_api/handler"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
)

func TestHandler(t *testing.T) {

	handler := rest.ResourceHandler{
		DisableJsonIndent: true,
		ErrorLogger:       log.New(ioutil.Discard, "", 0)}
	handler.SetRoutes(
		&rest.Route{"POST", "/api/coach/checkin", h.Checkin},
		&rest.Route{"POST", "/api/coach/checkout", h.Checkout},
	)

	recorded := test.RunRequest(t, &handler, test.MakeSimpleRequest(
		"POST", "http://www.sprint3r.com/api/coach/checkin",
		&map[string]string{"name": "iporsut", "league": "dtac"}))
	recorded.CodeIs(201)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{"id":"53f87e7ad18a68e0a884d31e"}`)

	recorded = test.RunRequest(t, &handler, test.MakeSimpleRequest(
		"POST", "http://www.sprint3r.com/api/coach/checkout",
		&map[string]string{"name": "iporsut", "league": "dtac"}))
	recorded.CodeIs(202)
	recorded.ContentTypeIsJson()
}
