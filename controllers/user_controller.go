package controllers

import (
	"complimentsapi/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetProfileHandler ... gets a user's profile
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	user, err := models.GetUserByUsername(username)
	if err != nil {
		fmt.Fprintln(w, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
