package dto

type PasswordResetRequest struct {
	Email    string `json:"email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
