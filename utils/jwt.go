package hash

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("my_secret_key")

func GenerateToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": id,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(JwtKey)
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JwtKey, nil
	})
	if err != nil {
		return 0, errors.New("invalid token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}
	//emailId := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
