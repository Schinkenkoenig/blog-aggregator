package users

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Schinkenkoenig/blog-aggregator/internal/database"
	"github.com/Schinkenkoenig/blog-aggregator/internal/helpers"
	"github.com/google/uuid"
)

type UsersController struct {
	DB *database.Queries
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

func FromDatabaseUser(user *database.User) *UserResponse {
	return &UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UserResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Id        uuid.UUID `json:"id"`
}

func (controller *UsersController) getUserByApiKeyHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := helpers.GetApiKey(r)
	if err != nil {
		helpers.RespondWithError(w, 400, "expecting api key in authorization header")
		return
	}

	fmt.Printf("apiKey: %s", apiKey)

	user, err := controller.DB.GetUserByApiKey(context.Background(), apiKey)
	if err != nil {
		helpers.RespondWithError(w, 404, "no matching user found")
		return
	}

	resp := FromDatabaseUser(&user)

	helpers.RespondWithJSON(w, 200, resp)
}

func (controller *UsersController) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// get data out of request

	req := helpers.NewRequestFromReader[CreateUserRequest](r.Body)
	if req == nil {
		helpers.RespondWithError(w, 400, "body could not be parsed")
		return
	}

	// act on database
	user, err := controller.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      req.Name,
	})
	if err != nil {
		helpers.RespondWithError(w, 500, "could not save user to db")
		return
	}

	// transform and return
	resp := FromDatabaseUser(&user)
	helpers.RespondWithJSON(w, 201, resp)
}
