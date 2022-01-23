package userController

import (
	"encoding/json"
	"net/http"
	"rest-api-go/database"
	userModel "rest-api-go/models"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var listUsers []userModel.User
	database.DB.Find(&listUsers)
	json.NewEncoder(w).Encode(listUsers)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user userModel.User
	database.DB.First(&user, id)
	json.NewEncoder(w).Encode(user)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var user userModel.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}
