package dto

type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binfding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId      uint64 `json:"user_id" form:"user_id" binding:"required"`
}

type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required" `
	Description string `json:"description" form:"description" binding:"required"`
	UserId      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
