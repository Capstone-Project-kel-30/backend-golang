package superadmin_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/mashbens/cps/business/superadmin"
	jwtService "github.com/mashbens/cps/business/user"

	sAdminEntity "github.com/mashbens/cps/business/superadmin/entity"
)

var service superadmin.SuperAdminService
var sAdmin1, sAdmin2 sAdminEntity.SuperAdmin

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAdminByID(t *testing.T) {
	t.Run("Expect to find super admin by id", func(t *testing.T) {
		sAdminID := int(sAdmin1.ID)
		sAdminIDs := strconv.Itoa(sAdminID)
		user, err := service.FindSuperAdminByID(sAdminIDs)
		if err != nil {
			t.Error("Expect to find user by id", err)
		} else {
			if user.ID != 1 {
				t.Errorf("Expected %d, got %d", 0, user.ID)
			}
		}
	})
}

func TestFindADminByName(t *testing.T) {
	t.Run("Expect to find super admin by name", func(t *testing.T) {
		sAdmin := sAdmin1.Name
		user, err := service.FindSuperAdminByName(sAdmin)
		if err != nil {
			t.Error("Expect to find super admin by name", err)
		} else {
			if user.Name != "John" {
				t.Errorf("Expected %s, got %s", "John", user.Name)
			}
		}
	})
}

func TestCreateAdmin(t *testing.T) {
	// exiting user
	t.Run("Expect to create super admin", func(t *testing.T) {
		_, err := service.CreateSuperAdmin(sAdmin2)
		if err == nil {
			t.Error("Expext error is not nil. Error: ", err)
			t.FailNow()
		} else if err != err {
			t.Error("Expet error super admin created. Error", err)
		}
	})
}

func TestLogin(t *testing.T) {
	t.Run("Expect to Login super admin ", func(t *testing.T) {
		_, err := service.Login(sAdmin2)
		if err == nil {
			t.Error("Expext error is invalid email or password. Error: ", err)
			t.FailNow()
		}
	})
}

func setup() {
	sAdmin1.ID = 1
	sAdmin1.Name = "John"
	sAdmin1.Password = "test123"

	sAdmin2.ID = 2
	sAdmin2.Name = "John"
	sAdmin2.Password = "test123"

	repo := newInMemoryRepository()
	jwtService := jwtService.NewJWTService()

	service = superadmin.NewSuperAdminService(repo, jwtService)
}

// fot exiting data
type inMemoryRepository struct {
	sAdminByID   map[string]sAdminEntity.SuperAdmin
	sAdminByName map[string]sAdminEntity.SuperAdmin
}

func newInMemoryRepository() *inMemoryRepository {
	var repo inMemoryRepository
	repo.sAdminByID = make(map[string]sAdminEntity.SuperAdmin)
	repo.sAdminByName = make(map[string]sAdminEntity.SuperAdmin)

	userID := int64(sAdmin1.ID)
	userIDs := strconv.FormatInt(userID, 10)
	repo.sAdminByID[userIDs] = sAdmin1
	repo.sAdminByName[sAdmin1.Name] = sAdmin1

	return &repo
}
func (r *inMemoryRepository) InsertSuperAdmin(sAdmin sAdminEntity.SuperAdmin) (sAdminEntity.SuperAdmin, error) {
	return sAdminEntity.SuperAdmin{}, nil
}

func (r *inMemoryRepository) FindSuperAdminByName(name string) (sAdminEntity.SuperAdmin, error) {
	return r.sAdminByName[name], nil
}

func (r *inMemoryRepository) FindSuperAdminByID(id string) (sAdminEntity.SuperAdmin, error) {
	return r.sAdminByID[id], nil
}

func (r *inMemoryRepository) Login(sAdmin sAdminEntity.SuperAdmin) (sAdminEntity.SuperAdmin, error) {
	return sAdmin, nil
}
