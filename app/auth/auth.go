package auth

/*
 * Package auth handles user authentication.
 */

import (
	"encoding/gob"
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	IsAdmin     bool   `json:"is_admin"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func init() {
	gob.Register(Authentication{})
	gob.Register(Token{})
}

func GenerateJWT(email string, isAdmin bool) (string, error) {
	var mySigningKey = []byte(app.Config.EncKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["is_admin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil

}
