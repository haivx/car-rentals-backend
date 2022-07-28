package services

import (
	"car-rentals-backend/model"
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var hmacSampleSecret []byte

func GenerateToken(userInfo *model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userInfo.Id,
		"username": userInfo.Username,
		"email":    userInfo.Email,
		"nbf":      time.Now().Add(time.Hour * 24 * 10).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

func ValidateToken(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method: %v")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil

	} else {
		return nil, err
	}
}
