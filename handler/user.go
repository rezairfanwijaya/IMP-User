package handler

import (
	"imp/helper"
	"imp/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerUser struct {
	userService user.IService
}

func NewUserHandler(userService user.IService) *handlerUser {
	return &handlerUser{userService: userService}
}

func (h *handlerUser) SignUp(c *gin.Context) {
	var input user.InputNewUser

	// bind
	if err := c.ShouldBindJSON(&input); err != nil {
		errBinding := helper.GenerateErrorBinding(err)
		response := helper.GenerateResponse(
			"gagal",
			http.StatusBadRequest,
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	userSignedUp, err, code := h.userService.CreateUser(input)
	if err != nil {
		response := helper.GenerateResponse(
			"gagal",
			code,
			err.Error(),
		)

		c.JSON(code, response)
		return
	}

	userFormatted := user.FormatUser(userSignedUp)

	response := helper.GenerateResponse(
		"sukses",
		code,
		userFormatted,
	)

	c.JSON(code, response)
}
