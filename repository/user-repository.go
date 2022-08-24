package repository

import (
	"fmt"
	"go_api_mysql_jwt_gin_gorm/entity"
	"go_api_mysql_jwt_gin_gorm/helper"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindById(ID uint64) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
}

type userConnection struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{db}
}

func (u *userConnection) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *userConnection) FindById(ID uint64) (entity.User, error) {
	var user entity.User
	err := u.db.Find(&user, ID).Error
	fmt.Println(err)
	if user.ID == 0 {
		return user, fmt.Errorf("no data from id %d", ID)
	}
	return user, err
}

func (u *userConnection) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := u.db.Where("email = ?", email).Find(&user).Error
	if user.ID == 0 {
		return user, fmt.Errorf("%s not found", email)
	}
	return user, err
}

func (u *userConnection) CreateUser(user entity.User) (entity.User, error) {
	hashed, err := helper.HasshAndSalt([]byte(user.Password))
	if err != nil {
		return user, err
	}
	user.Password = hashed

	err = u.db.Create(&user).Error
	return user, err
}

func (u *userConnection) UpdateUser(user entity.User) (entity.User, error) {
	if user.Password != "" {
		hashed, err := helper.HasshAndSalt([]byte(user.Password))
		if err != nil {
			return user, err
		}
		user.Password = hashed
	}

	err := u.db.Save(&user).Error
	return user, err
}
