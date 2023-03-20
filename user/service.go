package user

import (
	"fmt"
	"imp/helper"
	"net/http"
)

type IService interface {
	CreateUser(input InputNewUser) (User, error, int)
}

type service struct {
	userRepo IRepository
}

func NewService(userRepo IRepository) *service {
	return &service{userRepo}
}

func (s *service) CreateUser(input InputNewUser) (User, error, int) {
	// find user by username
	userByUsername, err := s.userRepo.FindByUsername(input.Username)
	if err != nil {
		return userByUsername, err, http.StatusInternalServerError
	}

	if userByUsername.ID != 0 {
		return userByUsername, fmt.Errorf(
			"username %v telah digunakan",
			input.Username,
		), http.StatusConflict
	}

	// save user to database
	var user User
	user.FullName = input.FullName
	user.Username = input.Username

	password, err := helper.HashingPassword(input.Password)
	if err != nil {
		return userByUsername, err, http.StatusInternalServerError
	}

	user.Password = password

	userSaved, err := s.userRepo.Save(user)
	if err != nil {
		return userByUsername, err, http.StatusInternalServerError
	}

	return userSaved, nil, http.StatusOK
}
