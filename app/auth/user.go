package auth

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/crypto"
	//"time"
)

type User struct {
	User_ID      int    `json:"user_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	IsAdmin      bool   `json:"is_admin"`
	IsVerified   bool   `json:"is_verified"`
	Role         string `json:"role"`
}

type CreateUserForm struct {
	Email    string
	Password string
}

func InsertUser(user *User) *User {
	execstring := "INSERT INTO users(email, password_hash, is_admin, is_verified) VALUES(?,?,?,?)"
	res, err := app.Config.Database.Exec(execstring, user.Email, user.PasswordHash, user.IsAdmin, user.IsVerified)
	if err != nil {
		panic(err) // TODO: proper error handling and recovery.
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err) // TODO: proper error handling and recovery.
	}

	fmt.Println("Inserted with ID ", int(id)) // TODO: Remove this when no longer needed.

	user.User_ID = int(id)

	return user
}

func CreateAndInsertUser(email string, password string, is_admin bool, is_verified bool) *User {
	user := &User{Email: email, IsAdmin: is_admin, IsVerified: is_verified}
	user.SetPassword(password)
	InsertUser(user) // TODO: should this be user = InsertUser(user) ?

	return user
}

func (user *User) SetPassword(password string) {
	user.PasswordHash = crypto.HashPassword(password)
}

func GetUserByEmail(email string) *User {
	user := new(User)

	sqlstatement := "SELECT user_id, email, password_hash, is_admin,  is_verified FROM users where email=?"

	row := app.Config.Database.QueryRow(sqlstatement, email)
	switch err := row.Scan(&user.User_ID, &user.Email, &user.PasswordHash, &user.IsAdmin, &user.IsVerified); err {
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
