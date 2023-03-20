package user

import (
	"fmt"
	"imp/helper"
	"net/http"
)

type IService interface {
	SignupUser(input InputNewUser) (User, int, error)
	LoginUser(input InputLoginUser) (User, int, error)
	GetUserByID(ID int) (User, int, error)
	GetAllTransaction(params ParamsGetAllUsers, url string) (PaginationUser, int, error)
}

type Service struct {
	userRepo IRepository
}

func NewService(userRepo IRepository) *Service {
	return &Service{userRepo}
}

func (s *Service) SignupUser(input InputNewUser) (User, int, error) {
	// find user by username
	userByUsername, err := s.userRepo.FindByUsername(input.Username)
	if err != nil {
		return userByUsername, http.StatusInternalServerError, err
	}

	if userByUsername.ID != 0 {
		return userByUsername, http.StatusConflict, fmt.Errorf(
			"username %v telah digunakan",
			input.Username,
		)
	}

	// save user to database
	var user User
	user.FullName = input.FullName
	user.Username = input.Username

	password, err := helper.HashingPassword(input.Password)
	if err != nil {
		return userByUsername, http.StatusInternalServerError, err
	}

	user.Password = password

	userSaved, err := s.userRepo.Save(user)
	if err != nil {
		return userByUsername, http.StatusInternalServerError, err
	}

	return userSaved, http.StatusOK, nil
}

func (s *Service) LoginUser(input InputLoginUser) (User, int, error) {
	// find user by username
	userByUsername, err := s.userRepo.FindByUsername(input.Username)
	if err != nil {
		return userByUsername, http.StatusInternalServerError, err
	}

	if userByUsername.ID == 0 {
		return userByUsername, http.StatusBadRequest, fmt.Errorf(
			"username %v belum terdaftar",
			input.Username,
		)
	}

	// verify password
	err = helper.VerifyPassword(
		userByUsername.Password,
		input.Password,
	)
	if err != nil {
		return userByUsername, http.StatusBadRequest, fmt.Errorf(
			"password salah",
		)
	}

	return userByUsername, http.StatusOK, nil
}

func (s *Service) GetUserByID(ID int) (User, int, error) {
	userByID, err := s.userRepo.FindByID(ID)
	if err != nil {
		return userByID, http.StatusInternalServerError, err
	}

	if userByID.ID == 0 {
		return userByID, http.StatusBadRequest, fmt.Errorf(
			"user dengan id %v tidak ditemukan",
			ID,
		)
	}

	return userByID, http.StatusOK, nil
}

func (s *Service) GetAllTransaction(params ParamsGetAllUsers, url string) (PaginationUser, int, error) {
	var paginationUser PaginationUser
	offset := params.Page * params.Limit

	users, totalData, totalPage, err := s.userRepo.FindAll(params, offset)
	if err != nil {
		return paginationUser, http.StatusInternalServerError, err
	}

	paginationUser.FirstPage = fmt.Sprintf(
		"%s?page=%v&order=%v&limit=%v",
		url,
		0,
		params.Order,
		params.Limit,
	)

	paginationUser.LastPage = fmt.Sprintf(
		"%s?page=%v&order=%v&limit=%v",
		url,
		totalPage,
		params.Order,
		params.Limit,
	)

	if params.Page > 0 {
		paginationUser.PreviousPage = fmt.Sprintf(
			"%s?page=%v&order=%v&limit=%v",
			url,
			params.Page-1,
			params.Order,
			params.Limit,
		)
	}

	if params.Page < totalPage {
		paginationUser.NextPage = fmt.Sprintf(
			"%s?page=%v&order=%v&limit=%v",
			url,
			params.Page+1,
			params.Order,
			params.Limit,
		)
	}

	if params.Page > totalPage {
		paginationUser.PreviousPage = ""
	}

	paginationUser.Page = params.Page
	paginationUser.Limit = params.Limit
	paginationUser.TotalData = totalData
	paginationUser.TotalPage = totalPage
	paginationUser.Users = users

	return paginationUser, http.StatusOK, nil
}
