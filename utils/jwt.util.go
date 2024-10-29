package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"example.com/event-booking/src/configs"
)

func GenerateToken(emailAddress string, customerId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"emailAddress": emailAddress,
		"customerId":   customerId,
		"exp":          time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(configs.Environment().APP_SECRET))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, error := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(configs.Environment().APP_SECRET), nil
	})
	if error != nil {
		return 0, error
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// emailAddress, _ := claims["emailAddress"].(string)
	customerId := int64(claims["customerId"].(float64))

	return customerId, nil
}
