package helper

import (
	"fmt"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

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
