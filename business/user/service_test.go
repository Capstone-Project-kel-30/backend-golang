package user_test

import (
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/mashbens/cps/business/user"
	userEntity "github.com/mashbens/cps/business/user/entity"
)

var service user.UserService

var authService user.AuthService
var jwtService user.JWTService
var user1, user2, user3 userEntity.User

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetCustomerByID(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		sID := strconv.Itoa(user1.ID)
		result, _ := service.FindUserByID(sID)
		if result.ID != user1.ID {
			t.Error("Expect found user id 1")
		}
	})
}

func TestGetCustomerByEmail(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		res, _ := service.FindUserByEmail(user1.Email)
		if res.ID != user1.ID {
			t.Error("Expect found user id 1")
		}

	})
}
func TestCreateUser(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		res, err := service.CreateUser(user1)
		if res != nil {
			t.Error("Expect create  user")
		} else if err == nil {
			t.Error("Expect create  user already exist", err)

		}

	})
}

func TestResetPass(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		res, err := service.ResetPassword(user1)
		if err != nil {
			t.Error("Expect reset user", err)
		} else if res == nil {
			t.Error("Expect reset user", err)
		}

	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		res, err := service.UpdateUser(user1)
		if err != nil {
			t.Error("Expect update user", err)
		} else if res == nil {
			t.Error("Expect update user", err)
		}

	})
}

func TestUpdateUserEx(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err := service.UpdateUserExpiry(strconv.Itoa(user1.ID), "12", "1")
		if err != nil {
			t.Error("Expect update user", err)
		}

	})
}

// ---------------auth------------

func TestGenerateToken(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err := jwtService.GenerateToken(strconv.Itoa(user1.ID))
		if err == "" {
			t.Error("Expect", err)
		}
	})
}

func TestGenerateTOTP(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err := authService.GenerateTOTP(strconv.Itoa(user1.ID))
		if err == "" {
			t.Error("Expect", err)
		}
	})
}
func TestSendOTPtoEmail(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err := authService.SendOTPtoEmail("091233", user1.Name, user1.Email)
		if err != nil {
			t.Error("Expect", err)
		}
	})
}
func TestSendEmailVerification(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		res, err := authService.SendEmailVerification(user1.Email)
		if err == nil {
			t.Error("Expect", err)
		}
		if res != nil {
			t.Error("Expect", err)

		}
	})
}
func TestSendEmailForgotPassword(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		res, err := authService.SendEmailForgotPassword(user1)
		if err != nil {
			t.Error("Expect", err)
		}
		if res == nil {
			t.Error("Expect", err)

		}
	})
}

func TestVerifyCredential(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err := authService.VerifyCredential(user1.Email, user1.Password)
		if err == nil {
			t.Error("Expect", err)
		}
	})
}

func TestLogin(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err, _ := authService.Login(user1)
		if err != nil {
			t.Error("Expect", err)
		}
	})
}

func TestRegister(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err, _ := authService.Register(user1)
		if err != nil {
			t.Error("Expect", err)
		}
	})
}

func TestResetPassword(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err, _ := authService.ResetPassword(user1)
		if err == nil {
			t.Error("Expect", err)
		}
	})
}

func setup() {
	user1.ID = 1
	user1.Name = "John"
	user1.Email = "test@mail.mail"
	user1.Password = "test123"
	user1.Phone = "0980932183"

	user2.ID = 2
	user2.Name = "Jane"
	user2.Email = "Jane@mail.mail"
	user2.Password = "test123"
	user2.Phone = "0980932183"

	repo := newInMemoryRepository()
	service = user.NewUserService(repo)
	jwtService = user.NewJWTService()
	authService = user.NewAuthService(service, jwtService)

}

type inMemoryRepository struct {
	UserID    map[int]userEntity.User
	UserEmail map[string]userEntity.User
	AllUser   []userEntity.User
}

func newInMemoryRepository() *inMemoryRepository {
	var repo inMemoryRepository
	repo.UserID = make(map[int]userEntity.User)
	repo.UserID[int(user1.ID)] = user1
	repo.UserID[int(user2.ID)] = user2

	repo.UserEmail = make(map[string]userEntity.User)
	repo.UserEmail[user1.Email] = user1
	repo.UserEmail[user2.Email] = user2

	return &repo
}
func (r *inMemoryRepository) InsertUser(user userEntity.User) (*userEntity.User, error) {

	return &user, nil
}

func (r *inMemoryRepository) FindByEmail(email string) (*userEntity.User, error) {
	data := r.UserEmail[email]
	if data.Email == "" {
		err := errors.New("data not found")
		return nil, err
	}
	return &data, nil
}

func (r *inMemoryRepository) FindByUserID(userID string) (*userEntity.User, error) {
	intID, _ := strconv.Atoi(userID)
	data := r.UserID[intID]
	if data.Email == "" {
		err := errors.New("data not found")
		return nil, err
	}
	return &data, nil
}

func (r *inMemoryRepository) UpdateUser(user userEntity.User) (*userEntity.User, error) {
	return &user, nil
}
func (r *inMemoryRepository) ResetPassword(user userEntity.User) (*userEntity.User, error) {
	return &user, nil
}
func (r *inMemoryRepository) UpdateUserExpiry(userID string, expiry string, memberType string) error {
	return nil
}
