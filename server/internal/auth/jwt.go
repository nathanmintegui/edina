package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	// TODO: add duration as env secret
	expiration := time.Second * time.Duration(28800)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
