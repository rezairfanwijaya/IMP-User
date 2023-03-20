package handler

import (
	"fmt"
	"imp/auth"
	"imp/helper"
	"imp/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerUser struct {
	userService user.IService
	authService auth.IAuth
}

func NewUserHandler(userService user.IService, authService auth.IAuth) *handlerUser {
	return &handlerUser{
		userService: userService,
		authService: authService,
	}
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
	userSignedUp, code, err := h.userService.SignupUser(input)
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

func (h *handlerUser) Login(c *gin.Context) {
	var input user.InputLoginUser

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
	userLoggedin, code, err := h.userService.LoginUser(input)
	if err != nil {
		response := helper.GenerateResponse(
			"gagal",
			code,
			err.Error(),
		)

		c.JSON(code, response)
		return
	}

	// genereate token
	token, err := h.authService.GenerateToken(userLoggedin.ID)
	if err != nil {
		response := helper.GenerateResponse(
			"gagal",
			http.StatusInternalServerError,
			err.Error(),
		)

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userFormatted := user.FormatUserLogin(userLoggedin, token)

	response := helper.GenerateResponse(
		"sukses",
		code,
		userFormatted,
	)

	c.JSON(code, response)
}

func (h *handlerUser) ListUser(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	order := c.DefaultQuery("order", "id desc")
	limit := c.DefaultQuery("limit", "10")

	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		response := helper.GenerateResponse(
			"gagal",
			http.StatusBadRequest,
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		response := helper.GenerateResponse(
			"gagal",
			http.StatusBadRequest,
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	params := user.ParamsGetAllUsers{
		Page:  pageNumber,
		Order: order,
		Limit: limitNumber,
	}

	path := c.Request.URL.Path
	url := fmt.Sprintf("http://localhost:3000%v", path)

	usersPagination, code, err := h.userService.GetAllTransaction(params, url)
	if err != nil {
		response := helper.GenerateResponse(
			"gagal",
			code,
			err.Error(),
		)

		c.JSON(code, response)
		return
	}

	response := helper.GenerateResponse(
		"sukses",
		code,
		usersPagination,
	)

	c.JSON(code, response)

}
