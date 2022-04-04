package config

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"goGinDemo/model/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres password=postgres dbname=public port=5432 sslmode=disable",
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

var jwtKey = []byte("OHM8Gf6C0QKBD05YSevmMypbUILJyY4lkSE1cxRjQ1bl0LKGprDc3q9Z+0yM8iRvFSWdJiErIzuJHVynZrPqZ/3+z/Vi2PC2UNxko7T/2WbF85M/cL96XcWID5e1o3ceGkxXZq6Ji1ghCLWQCpVSEx0faGhBtOdM6zY7YVdz4PAP9xcPk4+y2IVswv8CyVbzzzO+o5uO6RKTkmcT+bzigRYE1Y3BH3tY2Ff9Cg==")

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
	tokenString, err := token.SignedString(jwtKey)
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
		return jwtKey, nil
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
