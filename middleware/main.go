package middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"goGinDemo/config"
	"goGinDemo/model/users"
	"time"
)

type privateClaims struct {
	jwt.StandardClaims
	UID int `json:"uId"`
}

func GenerateToken(user usersRepo.User) (string, error) {
	expiresAt := time.Now().AddDate(0, 0, 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, privateClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expiresAt,
		},
		UID: user.ID,
	})
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (int, string, error) {
	var claims privateClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JwtKey, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}
	id := claims.UID
	username := claims.Subject
	return id, username, nil
}
