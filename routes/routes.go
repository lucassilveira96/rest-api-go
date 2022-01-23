package routes

import (
	"log"
	"net/http"
	userController "rest-api-go/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/user/", userController.GetUser).Methods("Get")
	r.HandleFunc("/user/{id}", userController.GetUserById).Methods("Get")
	r.HandleFunc("/user/", userController.Create).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
