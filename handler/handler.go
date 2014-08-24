package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func CheckoutHandler(w rest.ResponseWriter, r *rest.Request) {
	w.WriteHeader(http.StatusAccepted)
}

type Checkin struct {
	ID string `json:"id"`
}

func CheckinHandler(w http.ResponseWriter, r *http.Request) {
	c := Checkin{ID: "53f87e7ad18a68e0a884d31e"}
	b, _ := json.Marshal(c)

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
