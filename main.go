package main

import (
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

	userRepo := user.NewRepository(conn)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// endpoint
	router.POST("auth/signup", userHandler.SignUp)

	env, err := helper.GetENV(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := router.Run(env["DOMAIN"]); err != nil {
		log.Fatal(err.Error())
	}
}
