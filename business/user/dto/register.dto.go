package dto

type RegisterRequest struct {
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Phone    string `json:"phone" form:"phone" binding:"required,min=1"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
