package jwtToken

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	UserId string
	jwt.StandardClaims
}

func NewJWT(userId string, hmacSampleSecret []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		},
		UserId: userId,
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func ParseHmac(tokenString string, hmacSampleSecret []byte) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		fmt.Println(err)
		return err
	}

}
