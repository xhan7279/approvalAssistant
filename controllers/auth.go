package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Credentials to parse username/password
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims to be encoded to JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Authenticate to authorize user
func Authenticate(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	verified := verifyPassword(creds.Username, saltedPassword(creds.Password))

	// Wrong password || username
	if !verified {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Extend token for an hour
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
