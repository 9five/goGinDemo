package model

import (
	"goGinDemo/config"
	"goGinDemo/model/users"
)

var UR usersRepo.UsersRepository

func init() {
	db := config.NewDB()
	UR = usersRepo.NewUsersRepo(db)
}
