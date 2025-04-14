package handlers

import (
	"net/http"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
