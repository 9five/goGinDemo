package usersRepo

import (
	"gorm.io/gorm"
)

type UsersRepository interface {
	CreateUser(username string, password string) error
	GetUserByID(uid int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	UpdateUser(uid int, updateData User) error
	DeleteUser(uid int) error
}

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (ur *UsersRepo) CreateUser(username string, password string) error {
	user := User{
		Username: username,
		Password: password,
	}
	if res := ur.db.Create(user); res.Error != nil {
		return res.Error
	}
	return nil
}

func (ur *UsersRepo) GetUserByID(uid int) (*User, error) {
	var user User
	if res := ur.db.Where("id=?", uid).First(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (ur *UsersRepo) GetUserByUsername(username string) (*User, error) {
	var user User
	if res := ur.db.Where("username=?", username).First(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (ur *UsersRepo) UpdateUser(uid int, updateData User) error {
	if res := ur.db.Model(&User{}).Where("uid=?", uid).Updates(
		User{
			Username: updateData.Username,
			Password: updateData.Password,
		},
	); res.Error != nil {
		return res.Error
	}
	return nil
}

func (ur *UsersRepo) DeleteUser(uid int) error {
	if res := ur.db.Where("id=?", uid).Delete(&User{}); res.Error != nil {
		return res.Error
	}
	return nil
}
