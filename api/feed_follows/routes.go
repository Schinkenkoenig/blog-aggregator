package feedfollows

import "net/http"

func (ffc *FeedFollowsController) ApplyRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("POST /v1/feed_follows", ffc.followFeed)
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowId}", ffc.unfollowFeed)
	mux.HandleFunc("GET /v1/feed_follows", ffc.getAllOwnFeedFollows)

	return nil
}
