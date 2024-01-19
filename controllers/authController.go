package controllers

import (
	"encoding/json"
	"github.com/agunghasbi/schalter-api/models"
	u "github.com/agunghasbi/schalter-api/utils"
	"net/http"
	// "time"
)

var Register = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create() //Create user
	u.Respond(w, resp)
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Email, user.Password)

	// === Use this to set the client cookie for "token" as the JWT we just generated .Also set an expiry time which is the same as the token itself

	// token := resp["token"].(string)
	// expiresToken := resp["token_expires"].(time.Time)
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   token,
	// 	Expires: expiresToken,
	// })
	
	// === 

	u.Respond(w, resp)
}