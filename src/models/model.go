package models

import "time"

// User is for user details
type User struct {
	UserID    string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Age       int32     `json:"age"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

// Users is for collation users details in array
type Users = []User
