package service

import (
	"go_api_mysql_jwt_gin_gorm/dto"
	"go_api_mysql_jwt_gin_gorm/entity"
	"go_api_mysql_jwt_gin_gorm/repository"
)

type AuthService interface {
	FindAll() ([]entity.User, error)
	FindById(ID uint64) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	CreateUser(user dto.UserCreateDTO) (entity.User, error)
	// VerifyCredential(email string, password string) interface{}
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository}
}

func (s *authService) FindAll() ([]entity.User, error) {
	return s.userRepository.FindAll()
}

func (s *authService) FindById(ID uint64) (entity.User, error) {
	return s.userRepository.FindById(ID)
}

func (s *authService) FindByEmail(email string) (entity.User, error) {
	return s.userRepository.FindByEmail(email)
}

func (s *authService) CreateUser(user dto.UserCreateDTO) (entity.User, error) {
	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return s.userRepository.CreateUser(newUser)
}
