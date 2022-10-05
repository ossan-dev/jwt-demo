package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

// this key has to be kept secure. Whoever will have this key can authenticate our users
var secretKey = []byte("Abcd1234!!")

type tokenRes struct {
	Token string `json:"token"`
}

type tokenReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func generateToken(username, password string) (string, error) {
	// TODO: here you can add logic to check against a DB
	//...

	// create a new token by providing the cryptographic algorithm
	token := jwt.New(jwt.SigningMethodHS256)

	// set default/custom claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["username"] = username
	claims["password"] = password

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req tokenReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(err)
	}

	r.Body.Close()

	validator := validator.New()
	if err = validator.Struct(&req); err != nil {
		panic(err)
	}

	bearerToken, err := generateToken(req.Username, req.Password)
	if err != nil {
		panic(err)
	}

	res := &tokenRes{Token: bearerToken}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
}
