package class_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/mashbens/cps/business/class"
	classEntity "github.com/mashbens/cps/business/class/entity"
)

var service class.ClassService
var class1, class2 classEntity.Class

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	class1.ID = 1
	class1.Classname = "Zumba"
	class1.Description = "test123"
	class1.Trainer = "zoko"
	class1.Date = "17/12/2022"
	class1.ClassType = "online"
	class1.Capacity = 11
	class1.Status = "Available"
	class1.UserBooked = 2
	class1.Duration = 60

	repo := newInMemoryRepository()

	service = class.NewClassService(repo)
}

func TestFindMemberByID(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {

		err := service.FindAllClass("")
		if err != nil {
			t.Error("Expect", err)
		}
	})
}
func TestFindClassByID(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {
		classID := int(class1.ID)
		classIDs := strconv.Itoa(classID)
		res, err := service.FindClassByID(classIDs)
		if err != nil {
			t.Error("Expect", err)
		} else {
			if res.ID != 0 {
				t.Errorf("Expected %d, got %d", 0, res.ID)
			}
		}

	})
}
func TestFindAllClassOn(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {

		err := service.FindAllClassOn("")
		if err != nil {
			t.Error("Expect", err)
		}
	})
}
func TestFindAllClassOff(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {

		err := service.FindAllClassOff("")
		if err != nil {
			t.Error("Expect", err)
		}
	})
}
func TesFindClassByIDt(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {
		classID := int(class1.ID)
		classIDs := strconv.Itoa(classID)
		res, err := service.FindClassByID(classIDs)
		if err != nil {
			t.Error("Expect", err)
		} else {
			if res != nil {
				t.Error("Expect", err)

			}
		}
	})
}
func TestFindClassOffByID(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {

		classID := int(class1.ID)
		classIDs := strconv.Itoa(classID)
		res, err := service.FindClassOffByID(classIDs)
		if err != nil {
			t.Error("Expect", err)
		} else {
			if res == nil {
				t.Error("Expect", err)

			}
		}
	})
}
func TestFindClassOnByID(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {

		classID := int(class1.ID)
		classIDs := strconv.Itoa(classID)
		res, err := service.FindClassOnByID(classIDs)
		if err != nil {
			t.Error("Expect", err)
		} else {
			if res == nil {
				t.Error("Expect", err)

			}
		}
	})
}

// func TestInsertClass(t *testing.T) {
// 	t.Run("Expect", func(t *testing.T) {

// 		res, err := service.InsertClass(class1)
// 		if err == nil {
// 			t.Error("Expect", err)
// 		} else {
// 			if res != nil {
// 				t.Error("Expect", err)

// 			}
// 		}
// 	})
// }
// func TestUpdateClass(t *testing.T) {
// 	t.Run("Expect", func(t *testing.T) {

// 		res, err := service.UpdateClass(class1)
// 		if err != nil {
// 			t.Error("Expect", err)
// 		} else {
// 			if res != nil {
// 				t.Error("Expect", err)
// 			}
// 		}
// 	})
// }
func TestDeleteClass(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {
		classID := int(class1.ID)
		classIDs := strconv.Itoa(classID)
		err := service.DeleteClass(classIDs, classIDs)
		if err != nil {
			t.Error("Expect", err)
		}
	})
}
func TestUpdateClassStatus(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {
		classID := int(class1.ID)
		classIDs := strconv.Itoa(classID)
		err := service.UpdateClassStatus(classIDs, "Full Booked")
		if err != nil {
			t.Error("Expect", err)
		}
	})
}
func TestUpdateUserBooked(t *testing.T) {
	t.Run("Expect", func(t *testing.T) {

		err := service.UpdateUserBooked("")
		if err != nil {
			t.Error("Expect", err)
		}
	})
}

type inMemoryRepository struct {
	classID   map[int]classEntity.Class
	UserEmail map[string]classEntity.Class
	AllUser   []classEntity.Class
}

func newInMemoryRepository() *inMemoryRepository {
	var repo inMemoryRepository
	repo.classID = make(map[int]classEntity.Class)
	repo.classID[int(class1.ID)] = class1

	return &repo
}
func (r *inMemoryRepository) FindAllClass(search string) (data []classEntity.Class) {
	return
}
func (r *inMemoryRepository) FindClassByID(classID string) (classEntity.Class, error) {
	return classEntity.Class{}, nil

}
func (r *inMemoryRepository) FindAllClassOn(search string) (data []classEntity.Class) {
	return

}
func (r *inMemoryRepository) FindAllClassOff(search string) (data []classEntity.Class) {
	return
}
func (r *inMemoryRepository) FindClassOnByID(classID string) (classEntity.Class, error) {
	return classEntity.Class{}, nil

}
func (r *inMemoryRepository) FindClassOffByID(classID string) (classEntity.Class, error) {
	return classEntity.Class{}, nil

}
func (r *inMemoryRepository) InsertClass(class classEntity.Class) (classEntity.Class, error) {
	return class, nil
}
func (r *inMemoryRepository) UpdateClass(class classEntity.Class) (classEntity.Class, error) {
	return class, nil
}
func (r *inMemoryRepository) DeleteClass(classID string) error {
	return nil
}
func (r *inMemoryRepository) UpdateClassStatus(classID string, status string) error {
	return nil
}
func (r *inMemoryRepository) UpdateUserBooked(classID string, userBooked int) error {
	return nil
}
