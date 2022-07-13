package user_test

import (
	"errors"
	"os"
	"strconv"
	"strings"
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

func TestGetCustomerByID(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		sID := strconv.Itoa(user1.ID)
		result, _ := service.FindUserByID(sID)
		if result.ID != user1.ID {
			t.Error("Expect found customer id 1")
		}
	})
}

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
func TestGetCustomerByEmail(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		_, err := service.FindUserByEmail(user1.Email)
		if !ErrorContains(err, "") {
			t.Errorf("unexpected error: %v", err)
		}

	})
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
