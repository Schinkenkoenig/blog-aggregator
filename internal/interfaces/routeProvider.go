package interfaces

import "net/http"

type RouteProvider interface {
	ApplyRoutes(mux *http.ServeMux) error
}
