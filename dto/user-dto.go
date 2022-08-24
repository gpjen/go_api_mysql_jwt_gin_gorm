package dto

type UserCreateDTO struct {
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type UserUpdateDTO struct {
	Name     string `json:"name,omitempty" form:"name,omitempty"  validate:"omitempty,min=3"`
	Email    string `json:"email,omitempty" form:"email,omitempty"  validate:"omitempty,email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"omitempty,min=6"`
}

type UserLoginDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}
