package main

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	resp := ErrorResponse{Error: msg}
	data, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	data, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
