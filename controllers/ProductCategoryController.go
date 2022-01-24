package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-go/database"
	"rest-api-go/models"

	"github.com/gorilla/mux"
)

func GetProductCategory(w http.ResponseWriter, _ *http.Request) {
	var listProductsCategory []models.ProductCategory
	database.DB.Find(&listProductsCategory)
	json.NewEncoder(w).Encode(listProductsCategory)
}

func GetProductCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var productCategory models.ProductCategory
	database.DB.First(&productCategory, id)
	json.NewEncoder(w).Encode(productCategory)
}

func CreateProductCategory(w http.ResponseWriter, r *http.Request) {
	var productCategory models.ProductCategory
	json.NewDecoder(r.Body).Decode(&productCategory)
	database.DB.Create(&productCategory)
	json.NewEncoder(w).Encode(&productCategory)
}

func UpdateProductCategory(w http.ResponseWriter, r *http.Request) {
	var productCategory models.ProductCategory
	json.NewDecoder(r.Body).Decode(&productCategory)
	database.DB.Save(&productCategory)
	json.NewEncoder(w).Encode(&productCategory)
}

func DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var productCategory models.ProductCategory
	database.DB.Delete(&productCategory, id)
	json.NewEncoder(w).Encode(productCategory)
}
