package feeds

import "net/http"

func (fc *FeedsController) ApplyRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("POST /v1/feeds", fc.CreateFeedHandler)
	mux.HandleFunc("GET /v1/feeds", fc.GetAllFeeds)

	return nil
}
