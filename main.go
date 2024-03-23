package main

import (
	"database/sql"
	"fmt"
	"net/http"

	healthApi "github.com/Schinkenkoenig/blog-aggregator/api/health"
	"github.com/Schinkenkoenig/blog-aggregator/api/users"
	"github.com/Schinkenkoenig/blog-aggregator/internal/database"
	"github.com/Schinkenkoenig/blog-aggregator/internal/interfaces"
	_ "github.com/lib/pq"
)

func applyAllRoutes(m *http.ServeMux, providers ...interfaces.RouteProvider) {
	for _, r := range providers {
		r.ApplyRoutes(m)
	}
}

func main() {
	config, err := LoadConfiguration()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", config.ConnectionString)
	if err != nil {
		panic(err)
	}
	dbQueries := database.New(db)

	mux := http.NewServeMux()
	middlewared := middlewareCors(mux)

	userController := users.UsersController{DB: dbQueries}

	applyAllRoutes(mux,
		healthApi.HealthControllerRoutes{},
		&userController)

	err = http.ListenAndServe(fmt.Sprintf(":%s", config.Port), middlewared)

	panic(err)
}
