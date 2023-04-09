package handler

import (
	"encoding/json"
	"net/http"
)

type EchoHandler struct {
}

type EchoResponse struct {
	Data string `json:"data"`
}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (h *EchoHandler) Echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&EchoResponse{Data: "echo"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
