package store

import (
	"log"
	"time"
	"user-crud/models"
)

// for now I set static data, it's remove after I used the DB.
var Users = []models.User{
	{
		UserID:    "07",
		FirstName: "Mahendrasingh",
		LastName:  "Dhoni",
		Email:     "msd@gmail.com",
		Age:       35,
	},
}

// GetUsers retirve user details from db
func GetUsers() (models.Users, error) {

	return Users, nil
}

func AddUserDB(data models.User, userID string, time time.Time) (models.User, error) {
	result := data
	result.UserID = userID
	result.CreatedAt = time

	Users = append(Users, result)
	log.Println("INFO: store/store.go Record added successfully Inserted Data: ", result)
	return result, nil

}

func UpdateByUID(payload models.User, index int) error {
	payload.CreatedAt = Users[index].CreatedAt
	payload.UpdateAt = time.Now()
	Users[index] = payload
	log.Println("INFO: store/store.go : User update successful. ")
	return nil
}
