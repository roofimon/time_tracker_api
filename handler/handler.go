package handler

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func Checkin(w rest.ResponseWriter, r *rest.Request) {
	w.WriteHeader(http.StatusCreated)
	data := map[string]string{"id": "53f87e7ad18a68e0a884d31e"}
	w.WriteJson(data)
}

func Checkout(w rest.ResponseWriter, r *rest.Request) {
	w.WriteHeader(http.StatusAccepted)
}
