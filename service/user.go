/**
 * Package service provides the business logic for user management.
 * It interacts with the repository layer to perform operations such as login, registration, and password validations.
 *
 * The UserService interface defines the methods available for the user service:
 * - Login: Authenticates a user based on credentials.
 * - Register: Registers a new user into the system.
 * - CheckPassLength: Validates the length of the password.
 * - CheckPassAlphabet: Validates the characters used in the password.
 *
 * The userService struct implements the UserService interface and contains:
 * - userRepository: A repository layer object to interact with the user data.
 *
 * Functions:
 * - NewUserService: Constructor for creating a new instance of userService.
 * - Login: Implementation to authenticate a user.
 * - Register: Implementation to register a new user.
 * - CheckPassLength: Implementation to check if the password length is less than or equal to 5.
 * - CheckPassAlphabet: Implementation to check if the password contains only alphabetic characters.
 *
 * This service does not directly handle HTTP requests but is used by API handlers that define routes for user interactions.
 */

package service

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
)

type UserService interface {
	Login(user model.User) error
	Register(user model.User) error

	CheckPassLength(pass string) bool
	CheckPassAlphabet(pass string) bool
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Login(user model.User) error {
	return s.userRepository.CheckAvail(user)
}

func (s *userService) Register(user model.User) error {
	return s.userRepository.Add(user)
}

func (s *userService) CheckPassLength(pass string) bool {
	return len(pass) <= 5
}

func (s *userService) CheckPassAlphabet(pass string) bool {
	for _, char := range pass {
		if !('a' <= char && char <= 'z' || 'A' <= char && char <= 'Z') {
			return false
		}
	}
	return true
}
