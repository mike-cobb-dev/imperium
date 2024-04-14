package handlers

import "net/http"

type NotFoundHandler struct{}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// add templates here, for now test error
	w.Write([]byte("Error"))
}
