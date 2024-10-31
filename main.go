// main.go
package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"dating-app/controllers"
	"dating-app/database"
)

func main() {
	database.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/profiles", controllers.ViewProfiles).Methods("GET")
	r.HandleFunc("/swipe", controllers.Swipe).Methods("POST")
	r.HandleFunc("/purchase", controllers.PurchasePremium).Methods("POST")

	http.ListenAndServe(":8080", r)
}
