package sign

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken() (string, error) {
	privateKeyFile, err := os.Open("/home/ivan/Projects/go/golang-jwt/go-jws/sign_rsa/certs/id_rsa")
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

	token := jwt.New(jwt.SigningMethodRS256)

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

type CustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func ValidateToken(tokenSigned string) (res map[string]interface{}, err error) {
	publicKeyFile, err := os.Open("/home/ivan/Projects/go/golang-jwt/go-jws/sign_rsa/certs/id_rsa.pub")
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

	token, err := jwt.ParseWithClaims(tokenSigned, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		res = make(map[string]interface{}, 2)
		res["username"] = claims.Username
		res["password"] = claims.Password
		return
	}

	return nil, fmt.Errorf("token invalid: %v", tokenSigned)
}
