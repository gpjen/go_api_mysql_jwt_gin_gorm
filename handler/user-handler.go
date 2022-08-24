package handler

import (
	"fmt"
	"go_api_mysql_jwt_gin_gorm/dto"
	"go_api_mysql_jwt_gin_gorm/helper"
	"go_api_mysql_jwt_gin_gorm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (h *userHandler) CreateUser(c *gin.Context) {
	var newDataUser dto.UserCreateDTO
	err := c.ShouldBindJSON(&newDataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Error json request", err.Error()))
		return
	}

	v := validator.New()
	err = v.Struct(newDataUser)
	if err != nil {
		var errValidator []string

		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errValidator = append(errValidator, errMessage)
		}

		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed validate field request", errValidator))
		return
	}

	data, err := h.userService.CreateUser(newDataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed create new user", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseOK("create new user", data))
}
