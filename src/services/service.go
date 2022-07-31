package services

import (
	"errors"
	"log"
	"time"
	"user-crud/comman"
	"user-crud/models"
	"user-crud/store"

	"github.com/google/uuid"
)

const (
	ErrorGettingDataFromStore = "service/service.go error_in_retriving_store_data"
	ErrorDataNotFound         = "data_not_found"
)

// AllUsersDetails is for added business logic for fetch user details
func AllUsersDetails() (models.Users, error) {
	users, err := store.GetUsers()
	if err != nil {
		log.Println("error in store.GetUsers()", "actual error ", err)
		return models.Users{}, err
	}

	if len(users) == 0 {
		log.Println("error in service/service.go", ErrorDataNotFound)
		return users, errors.New(ErrorDataNotFound)

	}
	return users, nil
}

// AddUserLogic add single user
func AddUserLogic(payload models.User) (models.User, error) {

	userID := uuid.New().String()
	createTime := time.Now()

	result, err := store.AddUserDB(payload, userID, createTime)
	if err != nil {
		log.Println("Error: services/service.go ", err)
	}

	return result, nil
}

// UpdateByUserID is used for update any user by their id
func UpdateByUserID(data models.User, userID string) error {

	// Note: all logic will be change once we start using db

	index := getIndex(store.Users, userID)
	if index <= 0 {
		log.Println("Error: services/services.go", comman.ErrorNoDataFound)
		return errors.New(comman.ErrorNoDataFound)

	}

	data.UserID = userID
	if ok := store.UpdateByUID(data, index); ok != nil {
		log.Println("Error: services/services.go err: error_while_update_data")
		return errors.New("error_while_update_data")
	}

	return nil
}

// DeletUserbyID is used to delete single user by userid
func DeletUserbyID(userID string) error {
	index := getIndex(store.Users, userID)
	if index <= 0 {
		log.Println("Error: services/services.go", comman.ErrorNoDataFound)
		return errors.New(comman.ErrorNoDataFound)

	}
	// this logic will remove once we used DB
	store.Users = append(store.Users[:index], store.Users[index+1:]...)
	log.Println("INFO: user deleted successfuly.")

	return nil
}

// getIndex is used for find index from array if data is present
func getIndex(fromData models.Users, searchKey string) int {

	for i := 0; i < len(fromData); i++ {
		if fromData[i].UserID == searchKey {
			return i
		}

	}
	return -1
}
