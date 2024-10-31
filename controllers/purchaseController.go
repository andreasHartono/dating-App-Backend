package controllers

import (
	"encoding/json"
	"net/http"

	"dating-app/database"
	"dating-app/models"
)

func PurchasePremium(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	var user models.User

	database.DB.First(&user, userID)
	user.IsPremium = true
	database.DB.Save(&user)

	json.NewEncoder(w).Encode("Premium package activated")
}
