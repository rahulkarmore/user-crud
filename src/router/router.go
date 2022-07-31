package routers

import (
	"net/http"
	"user-crud/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	routerInstance := mux.NewRouter()

	//UserPrifix is for prifix for all path/router
	userPrifix := routerInstance.PathPrefix("/user-crud").Subrouter()

	// Creating the project related routes.
	userPrifix.HandleFunc("/get-users", handler.GetAllUsers).Methods(http.MethodGet)
	userPrifix.HandleFunc("/add-user", handler.AddUser).Methods(http.MethodPost)               //add single user
	userPrifix.HandleFunc("/update-user/{userID}", handler.UpdateUser).Methods(http.MethodPut) //update single user
	userPrifix.HandleFunc("/delete-user/{userID}", handler.DeleteUserByID).Methods(http.MethodDelete)

	return routerInstance
}
