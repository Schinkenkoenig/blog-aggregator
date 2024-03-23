package users

import "net/http"

func (ucr *UsersController) ApplyRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("POST /v1/users", ucr.createUserHandler)
	mux.HandleFunc("GET /v1/users", ucr.getUserByApiKeyHandler)

	return nil
}
