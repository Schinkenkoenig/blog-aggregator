package health

import "net/http"

type HealthControllerRoutes struct{}

func (hcr HealthControllerRoutes) ApplyRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("GET /v1/readiness", readinessController)
	mux.HandleFunc("GET /v1/err", errorController)

	return nil
}
