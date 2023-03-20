package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func GenerateResponse(status string, code int, data interface{}) *responseAPI {
	return &responseAPI{
		Meta: meta{
			Status: status,
			Code:   code,
		},
		Data: data,
	}
}

func GetENV(path string) (map[string]string, error) {
	env, err := godotenv.Read(path)
	if err != nil {
		return env, fmt.Errorf(
			"gagal mengambil env file : %v",
			err.Error(),
		)
	}

	return env, nil
}

func HashingPassword(rawPassoword string) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(rawPassoword), 10)
	if err != nil {
		return "", fmt.Errorf(
			"gagal hashing password :%v",
			err.Error(),
		)
	}

	return string(passwordHashed), nil
}

func VerifyPassword(hashedPassword, rawPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)); err != nil {
		return fmt.Errorf(
			"gagal verifikasi password : %v",
			err.Error(),
		)
	}

	return nil
}

func GenerateErrorBinding(err error) []string {
	var errBinding []string

	for _, e := range err.(validator.ValidationErrors) {
		errBinding = append(errBinding, e.Error())
	}

	return errBinding
}
