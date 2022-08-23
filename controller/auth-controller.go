package controller

import (
	"go_api_mysql_jwt_gin_gorm/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	//service
}

func NewAuthController() AuthController {
	return &authController{}
}

func (a *authController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, helper.ResponseOK("mantap boos", "no data"))
}

func (a *authController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, helper.ResponseFail("gagal tapi manatap", "no err"))
}
