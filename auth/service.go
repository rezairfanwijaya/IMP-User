package auth

import (
	"fmt"
	"imp/helper"

	"github.com/golang-jwt/jwt/v4"
)

type IAuth interface {
	GenerateToken(userID int) (string, error)
}

type auth struct{}

func NewAuth() *auth {
	return &auth{}
}

func (a *auth) GenerateToken(userID int) (string, error) {
	// payload
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	env, err := helper.GetENV("../.env")
	if err != nil {
		return "", err
	}

	// signed token
	tokenSigned, err := token.SignedString([]byte(env["SECRET_KEY"]))
	if err != nil {
		return tokenSigned, fmt.Errorf(
			"gagal melakukan signed token : %v",
			err.Error(),
		)
	}

	return tokenSigned, nil
}

func (a *auth) VerifyToken(token string) (*jwt.Token, error) {
	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token")
		}

		env, err := helper.GetENV("../.env")
		if err != nil {
			return env, err
		}

		return []byte(env["SECRET_KEY"]), nil
	})

	if err != nil {
		return tokenParsed, err
	}

	return tokenParsed, nil
}
