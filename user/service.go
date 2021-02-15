package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(inpur LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password_hash = string(passwordHash)
	user.Role = "User"
	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("gak ada usernya")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}
