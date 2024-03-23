package health

import (
	"net/http"

	helpers "github.com/Schinkenkoenig/blog-aggregator/internal/helpers"
)

type ReadinessResponse struct {
	Status string `json:"status"`
}

func readinessController(w http.ResponseWriter, _ *http.Request) {
	helpers.RespondWithJSON(w, 200, ReadinessResponse{Status: "ok"})
}
