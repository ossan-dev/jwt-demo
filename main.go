package main

import (
	"fmt"
	"net/http"

	"golangjwt/handlers"
)

func main() {
	http.HandleFunc("/home", handlers.HandlePage)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/reserved", handlers.IsAuthorized(handlers.ReservedPage))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(fmt.Sprintf("err while listening on port 8000, err: %v", err))
	}
}
