package api

import (
	"errors"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// UserService contains methods of the books service
type UserService interface {
	New(user NewUserRequest) error
	GetUsers() ([]UserRequest, error)
}

// UserRepository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
	CreateUser(NewUserRequest) error
	GetUsers() ([]UserRequest, error)
}

type userService struct {
	storage UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		storage: userRepo,
	}
}

func hashPassword(password string) (hashedPassword string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func checkPassword(hashedPassword string, password string) (isPasswordValid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}

	return true
}

func (u *userService) New(user NewUserRequest) error {
	// do some basic validations
	if user.Email == "" {
		return errors.New("user service - email required")
	}

	if user.Username == "" {
		return errors.New("user service - username required")
	}

	if user.Password == "" {
		return errors.New("user service - password required")
	}

	if user.FirstName == "" {
		return errors.New("user service - first name required")
	}

	if user.Surname == "" {
		return errors.New("user service - surname required")
	}

	// do some basic normalization
	user.Email = strings.TrimSpace(user.Email)

	// TO DO - hash password
	user.Password = hashPassword(user.Password) // hash password

	err := u.storage.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (u *userService) GetUsers() ([]UserRequest, error) {
	users, err := u.storage.GetUsers()

	if err != nil {
		return []UserRequest{}, err
	}

	return users, nil
}
