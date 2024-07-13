/**
 * Package repository provides the data access layer for managing user entities.
 * It interacts with the database layer to perform operations such as adding and checking user availability.
 *
 * The UserRepository interface defines the methods available for the user repository:
 * - Add: Adds a new user to the database.
 * - CheckAvail: Checks the availability of a user by username.
 *
 * The userRepository struct implements the UserRepository interface and contains:
 * - db: A reference to the gorm.DB object for database operations.
 *
 * Functions:
 * - NewUserRepo: Constructor for creating a new instance of userRepository.
 * - Add: Implementation to add a new user to the database.
 * - CheckAvail: Implementation to check the availability of a user by username.
 */

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
