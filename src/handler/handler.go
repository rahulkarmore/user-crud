package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"user-crud/comman"
	"user-crud/models"
	"user-crud/services"

	"github.com/gorilla/mux"
)

const (
	ErrorInGetAllUsers = "handler/handler.go error_while_fetching_users"
	ErrorEncodeData    = "Error encoding response object"
	ErrorDecodeData    = "Error decoidng response object"
)

// GetAllUsers is for fetch all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	allUsers, err := services.AllUsersDetails()
	if err != nil {

		if err.Error() == services.ErrorDataNotFound {
			http.Error(w, "Data Not Found", http.StatusNotFound)
			return
		}
		log.Println(ErrorInGetAllUsers, "function name: GetAllUsers", "actual error: ", err, allUsers)
		http.Error(w, "Error while retriving data", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(allUsers); err != nil {
		log.Println(ErrorEncodeData, "function name: GetAllUsers", "actual error: ", err, allUsers)
		http.Error(w, ErrorEncodeData, http.StatusInternalServerError)

	}
}

// AddUser is for add single user in db
func AddUser(w http.ResponseWriter, r *http.Request) {

	// extract data first
	userData := models.User{}

	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		log.Println("handler/handler.go Error :", ErrorDecodeData)
		http.Error(w, ErrorDecodeData, http.StatusBadRequest)
		return
	}

	//adding userData
	response, err := services.AddUserLogic(userData)

	if err != nil {
		log.Println("Error: handler.handler.go", err)
	}

	result, err := json.Marshal(&response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	// send response
	// future scope: need a comman functioanlity for this
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

// UpdateUse handler used for update user by userID
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	updatePayload := models.User{}
	userID := mux.Vars(r)[comman.UserID]

	err := json.NewDecoder(r.Body).Decode(&updatePayload)
	if err != nil {
		log.Println("handler/handler.go Error :", ErrorDecodeData)
		http.Error(w, ErrorDecodeData, http.StatusBadRequest)
		return
	}

	err = services.UpdateByUserID(updatePayload, userID)

	if err != nil {
		if err.Error() == comman.ErrorNoDataFound {
			log.Println("Error: handler/handler.go err : ", err)
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Println("Error: handler/handler.go err : ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&updatePayload); err != nil {
		log.Println("Error: handler/handler.go err : ", err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}

}

// DeleteUserByID is used for delete user by userID
// Note: once db added we will use soft delete method
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	userID := mux.Vars(r)[comman.UserID]

	err := services.DeletUserbyID(userID)

	if err != nil {

		log.Printf("\n Error : handler/handler.go not able to delete user due to err: %v, userID: %v", err, userID)
		http.Error(w, "Error not able to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
