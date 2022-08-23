package handler

import (
	"go_api_mysql_jwt_gin_gorm/dto"
	"go_api_mysql_jwt_gin_gorm/helper"
	"go_api_mysql_jwt_gin_gorm/service"
	"net/http"
	"strconv"

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

func (h *userHandler) FindById(c *gin.Context) {
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Find by id", err.Error()))
		return
	}

	data, err := h.userService.FindById(uint64(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Find by id", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseOK("Find by id", data))
}

func (h *userHandler) FindByEmail(c *gin.Context) {
	var emailRequest dto.EmailRequestDto
	err := c.ShouldBindJSON(&emailRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("find email", err.Error()))
		return
	}

	data, err := h.userService.FindByEmail(emailRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("find email", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseOK("find email", data))
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var newDataUser dto.UserCreateDTO
	err := c.ShouldBindJSON(&newDataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("create new user", err.Error()))
		return
	}
	data, err := h.userService.CreateUser(newDataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("create new user", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseOK("create new user", data))
}
