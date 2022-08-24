package repository

import (
	"fmt"
	"go_api_mysql_jwt_gin_gorm/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindById(ID uint64) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := u.db.Where("active = ?", true).Find(&users).Error
	return users, err
}

func (u *userRepository) FindById(ID uint64) (entity.User, error) {
	var user entity.User
	err := u.db.Where("active = ?", true).Find(&user, ID).Error
	if user.ID == 0 {
		return user, fmt.Errorf("no data from id %d", ID)
	}
	return user, err
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?, active = ?", email, true).Find(&user).Error
	if user.ID == 0 {
		return user, fmt.Errorf("%s not found", email)
	}
	return user, err
}

func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
