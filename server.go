package main

import (
	"go_api_mysql_jwt_gin_gorm/config"
	"go_api_mysql_jwt_gin_gorm/handler"
	"go_api_mysql_jwt_gin_gorm/repository"
	"go_api_mysql_jwt_gin_gorm/service"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	db = config.ConnectDB()

	// user
	userRepo    = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepo)
	userHandler = handler.NewUserHandler(userService)
)

func main() {
	defer config.CloseDB(db)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	group := router.Group("api/v1")

	group.POST("/login", userHandler.Login)

	group.GET("/users", userHandler.FindAll)
	group.GET("/user/:id", userHandler.FindById)
	group.POST("/user", userHandler.CreateUser)
	group.PATCH("/user/:id", userHandler.UpdateUser)
	group.DELETE("/user/:id", userHandler.SoftDelete)

	if err := router.Run(":" + port); err != nil {
		panic(err.Error())
	}
}
