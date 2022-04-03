package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

var JwtKey = []byte("OHM8Gf6C0QKBD05YSevmMypbUILJyY4lkSE1cxRjQ1bl0LKGprDc3q9Z+0yM8iRvFSWdJiErIzuJHVynZrPqZ/3+z/Vi2PC2UNxko7T/2WbF85M/cL96XcWID5e1o3ceGkxXZq6Ji1ghCLWQCpVSEx0faGhBtOdM6zY7YVdz4PAP9xcPk4+y2IVswv8CyVbzzzO+o5uO6RKTkmcT+bzigRYE1Y3BH3tY2Ff9Cg==")
