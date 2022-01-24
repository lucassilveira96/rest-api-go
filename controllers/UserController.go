package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-go/database"
	"rest-api-go/models"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, _ *http.Request) {
	var listUsers []models.User
	database.DB.Find(&listUsers)
	json.NewEncoder(w).Encode(listUsers)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
	database.DB.First(&user, id)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
	database.DB.Delete(&user, id)
	json.NewEncoder(w).Encode(user)
}
