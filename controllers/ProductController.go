package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-go/database"
	"rest-api-go/models"

	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, _ *http.Request) {
	var listProducts []models.Product
	database.DB.Find(&listProducts)
	json.NewEncoder(w).Encode(listProducts)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product
	database.DB.First(&product, id)
	database.DB.First(&product.ProductCategory, product.ProductCategoriesId)
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Create(&product)
	json.NewEncoder(w).Encode(&product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Save(&product)
	json.NewEncoder(w).Encode(&product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product
	database.DB.Delete(&product, id)
	json.NewEncoder(w).Encode(product)
}
