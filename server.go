package main

import (
	"go_api_mysql_jwt_gin_gorm/config"
	"go_api_mysql_jwt_gin_gorm/controller"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	db   = config.ConnectDB()
	auth = controller.NewAuthController()
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

	if err := router.Run(":" + port); err != nil {
		panic(err.Error())
	}
}
