package sign

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken() (string, error) {
	privateKeyFile, err := os.Open("/home/ivan/Projects/go/golang-jwt/go-jws/certs/privatekey.pem")
	if err != nil {
		panic(err)
	}
	defer privateKeyFile.Close()
	privateKeyBytes, err := io.ReadAll(privateKeyFile)
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	token := jwt.New(jwt.SigningMethodRS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["username"] = "test"
	claims["password"] = "test"

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}

func ValidateToken(tokenSigned string) (res map[string]interface{}, err error) {
	publicKeyFile, err := os.Open("/home/ivan/Projects/go/golang-jwt/go-jws/certs/publickey.pem")
	if err != nil {
		panic(err)
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := io.ReadAll(publicKeyFile)
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		panic(err)
	}

	token, err := jwt.Parse(tokenSigned, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		res = make(map[string]interface{}, 2)
		res["username"] = claims["username"]
		res["password"] = claims["password"]
		return
	}

	return nil, fmt.Errorf("token invalid: %v", tokenSigned)
}
