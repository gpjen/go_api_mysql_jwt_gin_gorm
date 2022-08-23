package handler

import (
	"go_api_mysql_jwt_gin_gorm/helper"
	"go_api_mysql_jwt_gin_gorm/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.AuthService
}

func NewUserHandler(userService service.AuthService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) FindAll(c *gin.Context) {
	data, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("find All users", err.Error()))
	}
	c.JSON(http.StatusOK, helper.ResponseOK("find All users", data))
}
