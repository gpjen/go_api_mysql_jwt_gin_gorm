package handler

import (
	"go_api_mysql_jwt_gin_gorm/dto"
	"go_api_mysql_jwt_gin_gorm/helper"
	"go_api_mysql_jwt_gin_gorm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Login(c *gin.Context) {
	// get json data
	var dataUser dto.UserLoginDTO
	err := c.ShouldBindJSON(&dataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed user login", err.Error()))
		return
	}

	// validate form/json login
	v := validator.New()
	err = v.Struct(dataUser)
	if err != nil {
		errValidator := helper.ConvertErrToSliceOfString(err)
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed user login", errValidator))
		return
	}

	// check data
	data, err := h.userService.LoginUser(dataUser.Email, dataUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed user login", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseOK("User login", data))
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

	// binding json
	var newDataUser dto.UserCreateDTO
	err := c.ShouldBindJSON(&newDataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Error json request", err.Error()))
		return
	}

	// validate json
	v := validator.New()
	err = v.Struct(newDataUser)
	if err != nil {
		errValidator := helper.ConvertErrToSliceOfString(err)
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed validate field request", errValidator))
		return
	}

	// create user
	data, err := h.userService.CreateUser(newDataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed create new user", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseOK("create new user", data))
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	// catch param id and convert to unit64
	idString := c.Param("id")
	idNumber, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed update user", err.Error()))
		return
	}

	// binding json body
	var newDataUpdate dto.UserUpdateDTO
	err = c.ShouldBindJSON(&newDataUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed update user", err.Error()))
		return
	}

	// handle empty data all field
	if newDataUpdate.Email == "" && newDataUpdate.Name == "" && newDataUpdate.Password == "" {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed update user", []string{"no data change"}))
		return
	}

	// validate field json
	v := validator.New()
	err = v.Struct(newDataUpdate)
	if err != nil {
		errValidator := helper.ConvertErrToSliceOfString(err)
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed update user", errValidator))
		return
	}

	// update new data
	data, err := h.userService.UpdateUser(newDataUpdate, uint64(idNumber))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed update user", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   idNumber,
		"data": data,
	})
}

func (h *userHandler) SoftDelete(c *gin.Context) {
	// get param id
	idString := c.Param("id")
	idNumber, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed delete user", err.Error()))
		return
	}

	data, err := h.userService.SoftDelete(uint64(idNumber))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFail("Failed delete user", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseOK("Delete user data", data))

}
