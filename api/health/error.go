package health

import (
	"net/http"

	helpers "github.com/Schinkenkoenig/blog-aggregator/internal/helpers"
)

func errorController(w http.ResponseWriter, _ *http.Request) {
	helpers.RespondWithError(w, 500, "Internal server error")
}
