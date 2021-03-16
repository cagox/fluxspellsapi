package models

/*
 * Package user contains the user model.
 */

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/crypto"
	"time"
)

//User is meant to hold user related information in the Database.
type User struct {
	ID           int       `json:"user_id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordhash"`
	IsAdmin      bool      `json:"isadmin"`
	LastUpdated  time.Time `json:"lastupdated"`
	TimeCreated  time.Time `json:"timecreated"`
}

//CreateUserForm is a struct to facilitate creating user objects.
type CreateUserForm struct {
	Email    string
	Password string
}

func InsertUser(user *User) *User {
	res, err := app.Config.Database.Exec("INSERT INTO users(email, passwordhash, isadmin) VALUES(?,?,?)", user.Email, user.PasswordHash, user.IsAdmin)
	if err != nil {
		panic(err) //Do something else with this later.
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err) //Do something else with this later.
	}

	fmt.Println("Inserted with ID ", int(id))

	user.ID = int(id)

	return user
}

func CreateAndInsertUser(email string, password string, isAdmin bool) *User {
	user := &User{Email: email, IsAdmin: isAdmin}
	user.SetPassword(password)
	InsertUser(user)

	return user
}

func (user *User) SetPassword(password string) {
	user.PasswordHash = crypto.HashPassword(password)
}

func GetUserByEmail(email string) *User {
	user := new(User)

	sqlStatement := "SELECT user_id, email, passwordhash, isadmin FROM users where email=?"

	row, err := app.Config.Database.QueryRow(sqlStatement, email)
	switch err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.IsAdmin); err {
	case sql.ErrNoRows:
		fmt.Println("User " + email + " does not exist.")
	case nil:
		return user
	default:
		panic(err)
	}
	return user
}

func (user *User) AuthenticateUser(password string) bool {
	return crypto.ComparePassword(password, user.PasswordHash)
}

func init() {
	gob.Register(User{})
}
