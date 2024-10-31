package controllers

import (
	"encoding/json"
	"net/http"

	"dating-app/database"
	"dating-app/models"
)

func ViewProfiles(w http.ResponseWriter, r *http.Request) {
	var profiles []models.User
	userID := r.Context().Value("userID").(uint)

	database.DB.Raw(`
        SELECT * FROM users WHERE id NOT IN (SELECT target_user_id FROM swipe_histories WHERE user_id = ?) LIMIT 10`, userID).Scan(&profiles)

	json.NewEncoder(w).Encode(profiles)
}

func Swipe(w http.ResponseWriter, r *http.Request) {
	var swipeData struct {
		TargetUserID uint `json:"target_user_id"`
		Liked        bool `json:"liked"`
	}
	json.NewDecoder(r.Body).Decode(&swipeData)

	userID := r.Context().Value("userID").(uint)
	database.DB.Create(&models.SwipeHistory{UserID: userID, TargetUserID: swipeData.TargetUserID, Liked: swipeData.Liked})
	w.WriteHeader(http.StatusCreated)
}
