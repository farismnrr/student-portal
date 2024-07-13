package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	return u.db.Create(&user).Error
}

func (u *userRepository) CheckAvail(user model.User) error {
	return u.db.Where("username = ?", user.Username).First(&model.User{}).Error
}
