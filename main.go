package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Schinkenkoenig/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	connStr := os.Getenv("CONN")
	fmt.Println(port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	dbQueries := database.New(db)

	mux := http.NewServeMux()
	middlewared := middlewareCors(mux)

	mux.HandleFunc("GET /v1/readiness", func(w http.ResponseWriter, _ *http.Request) {
		type ReadinessResponse struct {
			Status string `json:"status"`
		}
		respondWithJSON(w, 200, ReadinessResponse{Status: "ok"})
	})

	mux.HandleFunc("GET /v1/err", func(w http.ResponseWriter, _ *http.Request) {
		respondWithError(w, 500, "Internal server error")
	})

	mux.HandleFunc("POST /v1/users", func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Name string `json:"name"`
		}

		type Response struct {
			Id        uuid.UUID `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Name      string    `json:"name"`
		}

		decoder := json.NewDecoder(r.Body)
		var req Request
		err := decoder.Decode(&req)
		if err != nil {
			respondWithError(w, 400, "body could not be parsed")
			return
		}

		user, err := dbQueries.CreateUser(context.Background(), database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      req.Name,
		})
		if err != nil {
			respondWithError(w, 500, "could not save user to db")
			return
		}

		resp := Response{
			Id:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		respondWithJSON(w, 201, resp)
	})

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), middlewared)

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
