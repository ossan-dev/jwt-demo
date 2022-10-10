package main

import (
	"fmt"

	"gojws/sign"
)

func main() {
	tokenString, err := sign.GenerateToken()
	if err != nil {
		panic(err)
	}

	claims, err := sign.ValidateToken(tokenString)
	if err != nil {
		panic(err)
	}

	for k, v := range claims {
		fmt.Printf("key: %q - value: %q\n", k, v)
	}
}
