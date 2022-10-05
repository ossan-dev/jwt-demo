package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func IsAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if there is the token
		if r.Header["Authorization"] != nil {

			// parse the token
			token, err := jwt.Parse(strings.Replace(r.Header["Authorization"][0], "Bearer ", "", 1), func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])))
				}
				return secretKey, nil
			})
			// err while parsing the token
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("err while parsing the jwt token!"))
				return
			}

			// token valid
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Printf("username: %q\n", claims["username"])
				fmt.Printf("password: %q\n", claims["password"])
				endpoint(w, r)
				return
			}

		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("you must provide the jwt token!"))
	})
}
