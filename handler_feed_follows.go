package main

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/NunoMiguel111/rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedFolows, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follows")
		return
	}

	respondWithJSON(w, http.StatusCreated, feedFolows)
}
