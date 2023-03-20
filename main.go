package main

import (
	"imp/auth"
	"imp/database"
	"imp/handler"
	"imp/helper"
	"imp/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	router.GET("user/userlist", authMiddleware(*authService, *userService), userHandler.ListUser)

	env, err := helper.GetENV(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := router.Run(env["DOMAIN"]); err != nil {
		log.Fatal(err.Error())
	}
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get header authorization
		authHeader := c.GetHeader("Authorization")

		// must contain bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.GenerateResponse(
				"unauthorized",
				http.StatusUnauthorized,
				"akses ditolak",
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get jwt
		tokenString := ""
		headerSplit := strings.Split(authHeader, " ")
		if len(headerSplit) == 2 {
			tokenString = headerSplit[1]
		}

		// jwt validate
		token, err := authService.VerifyToken(tokenString)
		if err != nil {
			response := helper.GenerateResponse(
				"unauthorized",
				http.StatusUnauthorized,
				"token tidak valid",
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get payload
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.GenerateResponse(
				"unauthorized",
				http.StatusUnauthorized,
				"token tidak valid",
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get userid from payload
		userID := int(claim["user_id"].(float64))

		// get user by userid in token
		userByUserID, _, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.GenerateResponse(
				"unauthorized",
				http.StatusUnauthorized,
				err.Error(),
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set context
		c.Set("currentUser", userByUserID)
	}
}
