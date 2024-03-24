package feedfollows

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Schinkenkoenig/blog-aggregator/internal/database"
	"github.com/Schinkenkoenig/blog-aggregator/internal/helpers"
	"github.com/google/uuid"
)

type FeedFollowsController struct {
	DB *database.Queries
}

type FollowFeedRequest struct {
	FeedId uuid.UUID `json:"feed_id"`
}

type FeedFollowResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
	Id        uuid.UUID `json:"id"`
}

func NewFromDatabase(feedFollow database.FeedsUser) *FeedFollowResponse {
	return &FeedFollowResponse{
		Id:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserId:    feedFollow.UserID,
		FeedId:    feedFollow.FeedID,
	}
}

func (ffc *FeedFollowsController) getAllOwnFeedFollows(w http.ResponseWriter, r *http.Request) {
	// get user id from api key
	apiKey, err := helpers.GetApiKey(r)
	if err != nil {
		helpers.RespondWithError(w, 401, "user not authenticated")
		return
	}

	user, err := ffc.DB.GetUserByApiKey(context.Background(), apiKey)
	if err != nil {
		helpers.RespondWithError(w, 404, "user not found")
		return
	}

	feedFollows, err := ffc.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		helpers.RespondWithError(w, 500, "could not execute db query")
	}

	feedFollowResp := make([]FeedFollowResponse, 0, len(feedFollows))

	for _, f := range feedFollows {
		feedFollowResp = append(feedFollowResp, *NewFromDatabase(f))
	}

	helpers.RespondWithJSON(w, 200, feedFollowResp)
}

func (ffc *FeedFollowsController) unfollowFeed(w http.ResponseWriter, r *http.Request) {
	apiKey, err := helpers.GetApiKey(r)
	if err != nil {
		helpers.RespondWithError(w, 401, "user not authenticated")
		return
	}

	feedFollowId, err := helpers.GetUuidFromPath(r, "feedFollowId")
	if err != nil {
		helpers.RespondWithError(w, 400, "need feed follow id in path")
		return
	}

	user, err := ffc.DB.GetUserByApiKey(context.Background(), apiKey)
	if err != nil {
		helpers.RespondWithError(w, 404, "user not found")
		return
	}

	err = ffc.DB.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		ID:     *feedFollowId,
	})
	if err != nil {
		fmt.Printf("userid %s, feedFollowId %s\n", user.ID, *feedFollowId)
		helpers.RespondWithError(w, 500, "could not unfollow")
		return
	}

	w.WriteHeader(200)
	w.Write([]byte{})
}

func (ffc *FeedFollowsController) followFeed(w http.ResponseWriter, r *http.Request) {
	// get user from request
	apiKey, err := helpers.GetApiKey(r)
	if err != nil {
		helpers.RespondWithError(w, 401, "user not authenticated")
		return
	}

	user, err := ffc.DB.GetUserByApiKey(context.Background(), apiKey)
	if err != nil {
		helpers.RespondWithError(w, 404, "user not found")
		return
	}

	req := helpers.NewRequestFromReader[FollowFeedRequest](r.Body)

	if req == nil {
		helpers.RespondWithError(w, 400, "expected feed_id in json body")
		return
	}

	feedFollow, err := ffc.DB.FollowFeed(context.Background(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    req.FeedId,
		UserID:    user.ID,
	})
	if err != nil {
		helpers.RespondWithError(w, 500, "could not follow feed")
	}

	resp := NewFromDatabase(feedFollow)

	helpers.RespondWithJSON(w, 201, resp)
}
