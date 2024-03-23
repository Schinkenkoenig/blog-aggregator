package users

import "net/http"

func (ucr *UsersController) ApplyRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("POST /v1/users", ucr.createUserHandler)

	return nil
}
