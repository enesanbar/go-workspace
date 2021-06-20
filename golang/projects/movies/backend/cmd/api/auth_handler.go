package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"

	"github.com/enesanbar/workspace/golang/projects/movies/backend/models"
	"golang.org/x/crypto/bcrypt"
)

// validUser is a dummy user
var validUser = models.User{
	ID:       10,
	Email:    "me@here.com",
	Password: "$2a$12$DEs0zVIDgOV1BNKu.e0TP.XP0r1c0Eaq/mwjAkScO/tmZ.Qhp4l3y", // hash for password "password"
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (a *application) Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		a.errorJSON(w, errors.New("unauthorised"))
		return
	}

	// simulate db with dummy user
	hashedPassword := validUser.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		a.errorJSON(w, errors.New("unauthorised"))
		return
	}

	claims := jwt.Claims{
		Registered: jwt.Registered{
			Issuer:    "mydomain.com",
			Subject:   fmt.Sprint(validUser.ID),
			Audiences: []string{"mydomain.com"},
			Expires:   jwt.NewNumericTime(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericTime(time.Now()),
			Issued:    jwt.NewNumericTime(time.Now()),
		},
	}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(a.config.jwt.secret))
	if err != nil {
		a.errorJSON(w, errors.New("error signing jwt"))
		return
	}

	a.writeJSON(w, http.StatusOK, string(jwtBytes), "response")

}
