package handler

import (
	"encoding/json"
	"net/http"
)

type Checkin struct {
	ID string `json:"id"`
}

func CheckinHandler(w http.ResponseWriter, r *http.Request) {
	c := Checkin{ID: "53f87e7ad18a68e0a884d31e"}
	b, _ := json.Marshal(c)

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
