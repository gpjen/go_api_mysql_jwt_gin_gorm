package service

import (
	"fmt"
	"go_api_mysql_jwt_gin_gorm/dto"
	"go_api_mysql_jwt_gin_gorm/entity"
	"go_api_mysql_jwt_gin_gorm/helper"
	"go_api_mysql_jwt_gin_gorm/repository"
	"os"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

type UserService interface {
	FindAll() ([]entity.User, error)
	FindById(ID uint64) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	LoginUser(email string, pwd string) (entity.User, error)
	CreateUser(user dto.UserCreateDTO) (entity.User, error)
	UpdateUser(user dto.UserUpdateDTO, ID uint64) (entity.User, error)
	SoftDelete(ID uint64) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

// user login
func (s *userService) LoginUser(email string, pwd string) (entity.User, error) {
	// check email user
	data, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return data, fmt.Errorf("email and password doesnt match")
	}

	// check password user
	matching, _ := helper.ComparePasword(pwd, data.Password)
	if !matching {
		return data, fmt.Errorf("email and password doesnt match")
	}

	// generate jwt token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = data.ID
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRETKEY")))
	if err != nil {
		return entity.User{}, fmt.Errorf("failed generate token")
	}
	data.Token = t

	return data, nil
}

// find all users
func (s *userService) FindAll() ([]entity.User, error) {
	return s.userRepository.FindAll()
}

// find user by ID
func (s *userService) FindById(ID uint64) (entity.User, error) {
	return s.userRepository.FindById(ID)
}

// find user by email
func (s *userService) FindByEmail(email string) (entity.User, error) {
	return s.userRepository.FindByEmail(email)
}

// create new user
func (s *userService) CreateUser(user dto.UserCreateDTO) (entity.User, error) {
	hash, err := helper.HasshAndSalt([]byte(user.Password))
	if err != nil {
		return entity.User{}, err
	}

	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hash,
	}

	return s.userRepository.CreateUser(newUser)
}

// update new User
func (s *userService) UpdateUser(user dto.UserUpdateDTO, ID uint64) (entity.User, error) {
	findData, err := s.userRepository.FindById(ID)
	if err != nil {
		return findData, err
	}

	if user.Name != "" {
		findData.Name = user.Name
	}
	if user.Email != "" {
		findData.Email = user.Email
	}
	if user.Password != "" {
		findData.Password, _ = helper.HasshAndSalt([]byte(user.Password))
	}

	return s.userRepository.UpdateUser(findData)
}

// soft delete user
func (s *userService) SoftDelete(ID uint64) (entity.User, error) {
	user, err := s.userRepository.FindById(ID)
	if err != nil {
		return user, err
	}

	// update active and save
	user.Active = false
	s.userRepository.UpdateUser(user)

	return user, nil
}
