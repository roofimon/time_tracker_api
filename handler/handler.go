package handler

import (
	"encoding/json"
	"net/http"
)

type Checkin struct {
	ID string `json:"id"`
}

type User struct {
	Name   string `json:"name"`
	League string `json:"league"`
}

type Tracker interface {
	CheckIn(user User) Checkin
}

type TimeTrackerHandler struct {
	Tracker
}

func (t *TimeTrackerHandler) Checkin(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	c := t.CheckIn(user)
	b, _ := json.Marshal(c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
