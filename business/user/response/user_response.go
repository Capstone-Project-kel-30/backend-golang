package response

import "gym-app/business/user/entity"

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Token string `json:"token,omitempty"`
	Totp  string `json:"totp,omitempty"`
}

func NewUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Phone: user.Phone,
	}
}
