package main

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/NunoMiguel111/rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	respondWithJSON(w, http.StatusCreated, feed)
}

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := cfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get all feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}
