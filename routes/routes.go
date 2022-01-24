package routes

import (
	"log"
	"net/http"
	"rest-api-go/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	//routes user
	r.HandleFunc("/user/", controllers.GetUser).Methods("Get")
	r.HandleFunc("/user/{id}", controllers.GetUserById).Methods("Get")
	r.HandleFunc("/user/", controllers.CreateUser).Methods("Post")
	r.HandleFunc("/user/", controllers.UpdateUser).Methods("Put")
	r.HandleFunc("/product/{id}", controllers.DeleteUser).Methods("Delete")

	//routes product
	r.HandleFunc("/product/", controllers.GetProduct).Methods("Get")
	r.HandleFunc("/product/{id}", controllers.GetProductById).Methods("Get")
	r.HandleFunc("/product/", controllers.CreateProduct).Methods("Post")
	r.HandleFunc("/product/", controllers.UpdateProduct).Methods("Put")
	r.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("Delete")

	//routes productcategory
	r.HandleFunc("/product/category/", controllers.GetProductCategory).Methods("Get")
	r.HandleFunc("/product/category/{id}", controllers.GetProductCategoryById).Methods("Get")
	r.HandleFunc("/product/category/", controllers.CreateProductCategory).Methods("Post")
	r.HandleFunc("/product/category/", controllers.UpdateProductCategory).Methods("Put")
	r.HandleFunc("/product/category/{id}", controllers.DeleteProductCategory).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8000", r))
}
