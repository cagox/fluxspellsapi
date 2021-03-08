package user

/*
 * Package user contains the user model.
 */


import (
"encoding/gob"
"time"
)

var Version = 1

//User is meant to hold user related information in the Database.
type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordhash"`
	IsAdmin      bool      `json:"isadmin"`
	Timestamp    time.Time `json:"timestamp"`
	LastUpdated  time.Time `json:"lastupdated"`
}

//CreateUserForm is a struct to facilitate creating user objects.
type CreateUserForm struct {
	Email    string
	Password string
}

func init() {
	gob.Register(User{})
}
