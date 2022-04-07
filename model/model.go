package model

import (
	"goGinDemo/config"
	"goGinDemo/model/users"
)

var UsersR usersRepo.UsersRepository

func init() {
	db := config.NewDB()
	UsersR = usersRepo.NewUsersRepo(db)
}
