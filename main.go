package main

import (
	"imp/auth"
	"imp/database"
	"imp/handler"
	"imp/helper"
	"imp/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// connection to db
	conn, err := database.NewConnection(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	authService := auth.NewAuth()

	userRepo := user.NewRepository(conn)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService, authService)

	// endpoint
	router.POST("auth/signup", userHandler.SignUp)
	router.POST("auth/login", userHandler.Login)
	router.GET("user/userlist", userHandler.ListUser)

	env, err := helper.GetENV(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := router.Run(env["DOMAIN"]); err != nil {
		log.Fatal(err.Error())
	}
}
