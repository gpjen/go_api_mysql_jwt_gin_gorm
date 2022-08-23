package main

import (
	"go_api_mysql_jwt_gin_gorm/config"
	"go_api_mysql_jwt_gin_gorm/controller"
	"go_api_mysql_jwt_gin_gorm/handler"
	"go_api_mysql_jwt_gin_gorm/repository"
	"go_api_mysql_jwt_gin_gorm/service"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	db   = config.ConnectDB()
	auth = controller.NewAuthController()

	userRepo    = repository.NewUserRepository(db)
	userService = service.NewAuthService(userRepo)

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

	group.GET("/login", auth.Login)
	group.GET("/register", auth.Register)
	group.GET("/users", userHandler.FindAll)
	group.GET("/user/:id", userHandler.FindById)
	group.POST("/email", userHandler.FindByEmail)
	group.POST("/user", userHandler.CreateUser)

	if err := router.Run(":" + port); err != nil {
		panic(err.Error())
	}
}
