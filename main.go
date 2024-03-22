package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	middlewared := middlewareCors(mux)

	mux.HandleFunc("GET /v1/readiness", func(w http.ResponseWriter, r *http.Request) {
		type ReadinessResponse struct {
			Status string `json:"status"`
		}
		respondWithJSON(w, 200, ReadinessResponse{Status: "ok"})
	})

	mux.HandleFunc("GET /v1/err", func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, 500, "Internal server error")
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), middlewared)

	panic(err)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
