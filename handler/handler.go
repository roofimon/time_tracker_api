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
	CheckOut(user User) error
}

type TimeTrackerHandler struct {
	Tracker
}

func (t *TimeTrackerHandler) checkin(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	c := t.CheckIn(user)
	b, _ := json.Marshal(c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (t *TimeTrackerHandler) checkout(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	t.CheckOut(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (t *TimeTrackerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.checkin(w, r)
}
