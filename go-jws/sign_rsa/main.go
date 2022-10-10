package main

import (
	"fmt"

	"signrsa/sign"
)

func main() {
	signedToken, err := sign.GenerateToken()
	if err != nil {
		panic(err)
	}

	claims, err := sign.ValidateToken(signedToken)
	if err != nil {
		panic(err)
	}

	for k, v := range claims {
		fmt.Printf("key: %q - value: %q\n", k, v)
	}
}
