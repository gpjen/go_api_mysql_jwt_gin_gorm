package dto

type EmailRequestDto struct {
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
}
