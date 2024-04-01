package feeds

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Schinkenkoenig/blog-aggregator/internal/database"
	"github.com/Schinkenkoenig/blog-aggregator/internal/helpers"
	"github.com/google/uuid"
)

type FeedsController struct {
	DB *database.Queries
}

type CreateFeedRequest struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type FeedResponse struct {
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserId        uuid.UUID  `json:"user_id"`
	Id            uuid.UUID  `json:"id"`
}

type ListFeedResponse []FeedResponse

func NewFromDatabase(feed database.Feed) *FeedResponse {
	var lastFetched *time.Time
	if feed.LastFetchedAt.Valid {
		lastFetched = &feed.LastFetchedAt.Time
	}

	return &FeedResponse{
		Id:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		LastFetchedAt: lastFetched,
		Name:          feed.Name,
		Url:           feed.Name,
		UserId:        feed.UserID,
	}
}

func (fc *FeedsController) GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := fc.DB.GetAllFeeds(context.Background())
	if err != nil {
		helpers.RespondWithError(w, 500, "could not load feeds")
		return
	}

	feedsResponse := make([]FeedResponse, 0, len(feeds))

	for _, f := range feeds {
		feedsResponse = append(feedsResponse, *NewFromDatabase(f))
	}

	helpers.RespondWithJSON(w, 200, feedsResponse)
}

func (fc *FeedsController) CreateFeedHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := helpers.GetApiKey(r)
	if err != nil {
		helpers.RespondWithError(w, 401, "not authenticated")
		return
	}

	req := helpers.NewRequestFromReader[CreateFeedRequest](r.Body)

	if req == nil {
		helpers.RespondWithError(w, 400, "bad request json")
		return
	}

	_, err = url.Parse(req.Url)
	if err != nil {
		helpers.RespondWithError(w, 400, "expected property url to be in url format")
		return
	}

	user, err := fc.DB.GetUserByApiKey(context.Background(), apiKey)
	if err != nil {
		helpers.RespondWithError(w, 404, "user not found")
		return
	}

	feed, err := fc.DB.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Url:       req.Url,
			Name:      req.Name,
			UserID:    user.ID,
		})
	if err != nil {
		helpers.RespondWithError(w, 500, "could not save feed")
		return
	}

	_, err = fc.DB.FollowFeed(context.Background(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		fmt.Println(err)
	}

	resp := NewFromDatabase(feed)

	helpers.RespondWithJSON(w, 201, resp)
}
