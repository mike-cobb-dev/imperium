package handlers

import (
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// template and other stuff to get and tie into the home request like user context and stuff

	w.Write([]byte("Homepage"))
}
