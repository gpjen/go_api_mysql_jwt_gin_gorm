package dto

type LoginDTO struct {
	Email    string `json:"eamil" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
