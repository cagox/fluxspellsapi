package api

import (
	"encoding/json"
	"github.com/cagox/fluxspellsapi/app/auth"
	"github.com/cagox/fluxspellsapi/app/models"
	"net/http"
)

func userLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.GetUserByEmail(email)

	validated := user.AuthenticateUser(password)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if validated {

		validToken, err := auth.GenerateJWT(user.Email, user.IsAdmin)
		if err != nil {
			json.NewEncoder(w).Encode("Failed to generate token")
			return
		}

		var token auth.Token
		token.Email = user.Email
		token.IsAdmin = user.IsAdmin
		token.TokenString = validToken
		json.NewEncoder(w).Encode(token)
	} else {
		err := "Invalid Email Address or Password"
		json.NewEncoder(w).Encode(err)
		return
	}

}
