package dto

type UserCreateDTO struct {
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name"  validate:"required,min=3"`
	Email    string `json:"email" form:"email"  validate:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min=6"`
}

type UserLoginDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}
