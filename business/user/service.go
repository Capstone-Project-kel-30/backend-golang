package user

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/mashbens/cps/business/user/entity"
)

type UserRepository interface {
	InsertUser(user entity.User) (*entity.User, error)
	UpdateUser(user entity.User) (*entity.User, error)
	ResetPassword(user entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByUserID(userID string) (*entity.User, error)
	UpdateUserExpiry(userID string, expiry string, memberType string) error
}

type UserService interface {
	CreateUser(user entity.User) (*entity.User, error)
	UpdateUser(user entity.User) (*entity.User, error)
	ResetPassword(user entity.User) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	FindUserByID(userID string) (*entity.User, error)
	UpdateUserExpiry(userID string, expiry string, memberType string) error
}

type userService struct {
	userRepo UserRepository
	validate *validator.Validate
}

func NewUserService(
	userRepo UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		validate: validator.New(),
	}
}

func (c *userService) CreateUser(user entity.User) (*entity.User, error) {

	_, err := c.userRepo.FindByEmail(user.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	usr, _ := c.userRepo.InsertUser(user)

	return usr, nil
}

func (c *userService) FindUserByEmail(email string) (*entity.User, error) {
	usr, _ := c.userRepo.FindByEmail(email)

	return usr, nil
}

func (c *userService) ResetPassword(user entity.User) (*entity.User, error) {

	u, _ := c.userRepo.ResetPassword(user)

	return u, nil
}

func (c *userService) FindUserByID(userID string) (*entity.User, error) {
	return c.userRepo.FindByUserID(userID)
}

func (c *userService) UpdateUser(user entity.User) (*entity.User, error) {
	usr, _ := c.userRepo.UpdateUser(user)

	return usr, nil
}

func (c *userService) UpdateUserExpiry(userID string, expiry string, memberType string) error {

	user := c.userRepo.UpdateUserExpiry(userID, expiry, memberType)
	return user
}
