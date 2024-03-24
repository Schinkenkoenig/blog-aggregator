package main

import (
	"database/sql"
	"fmt"
	"net/http"

	feedfollows "github.com/Schinkenkoenig/blog-aggregator/api/feed_follows"
	"github.com/Schinkenkoenig/blog-aggregator/api/feeds"
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

	// set up database queries
	db, err := sql.Open("postgres", config.ConnectionString)
	if err != nil {
		panic(err)
	}
	dbQueries := database.New(db)

	// setup http server, apply middlewares and routes
	mux := http.NewServeMux()
	middlewared := middlewareCors(mux)

	// create controller structs with injected services
	userController := users.UsersController{DB: dbQueries}
	feedController := feeds.FeedsController{DB: dbQueries}
	feedFollowsController := feedfollows.FeedFollowsController{DB: dbQueries}

	applyAllRoutes(mux,
		healthApi.HealthControllerRoutes{},
		&userController,
		&feedController,
		&feedFollowsController)

	err = http.ListenAndServe(fmt.Sprintf(":%s", config.Port), middlewared)

	panic(err)
}
