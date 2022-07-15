package newsletter_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/mashbens/cps/business/newsletter"
	newsletterEntity "github.com/mashbens/cps/business/newsletter/entity"
)

var service newsletter.NewsService
var news1, class2 newsletterEntity.News

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	t.Run("Expect id", func(t *testing.T) {

		err := service.FindAllNews("")
		if err != nil {
			t.Error("Expect id", err)
		}
	})
}
func TestFindNewsByID(t *testing.T) {
	t.Run("Expect id", func(t *testing.T) {
		news := int(news1.ID)
		n := strconv.Itoa(news)

		res, err := service.FindNewsByID(n)
		if err != nil {
			t.Error("Expect id", err)
		} else {
			if res.ID != 0 {
				t.Errorf("Expected %d, got %d", 0, res.ID)
			}
		}
	})
}

// func TestInsertNews(t *testing.T) {
// 	t.Run("Expect id", func(t *testing.T) {

// 		res, err := service.InsertNews(news1)
// 		if err != nil {
// 			t.Error("Expect id", err)
// 		} else {
// 			if res.ID != 0 {
// 				t.Errorf("Expected %d, got %d", 0, res.ID)
// 			}
// 		}
// 	})
// }
// func TestUpdateNews(t *testing.T) {
// 	t.Run("Expect id", func(t *testing.T) {

// 		res, err := service.UpdateNews(news1)
// 		if err != nil {
// 			t.Error("Expect id", err)
// 		} else {
// 			if res.ID != 0 {
// 				t.Errorf("Expected %d, got %d", 0, res.ID)
// 			}
// 		}
// 	})
// }
func TestDeleteNews(t *testing.T) {
	news := int(news1.ID)
	n := strconv.Itoa(news)

	err := service.DeleteNews(n, n)
	if err != nil {
		t.Error("Expect to find member by id", err)
	}
}

func setup() {
	news1.ID = 1
	news1.Title = "Zumba"
	news1.Date = "test123"
	news1.Content = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

	repo := newInMemoryRepository()

	service = newsletter.NewNewsService(repo)
}

type inMemoryRepository struct {
	newsID    map[int]newsletterEntity.News
	UserEmail map[string]newsletterEntity.News
	AllUser   []newsletterEntity.News
}

func newInMemoryRepository() *inMemoryRepository {
	var repo inMemoryRepository
	repo.newsID = make(map[int]newsletterEntity.News)
	repo.newsID[int(news1.ID)] = news1

	return &repo
}
func (r *inMemoryRepository) FindAllNews(search string) (data []newsletterEntity.News) {
	return
}
func (r *inMemoryRepository) FindNewsByID(newsID string) (newsletterEntity.News, error) {
	return newsletterEntity.News{}, nil
}
func (r *inMemoryRepository) InsertNews(news newsletterEntity.News) (newsletterEntity.News, error) {
	return newsletterEntity.News{}, nil
}
func (r *inMemoryRepository) UpdateNews(news newsletterEntity.News) (newsletterEntity.News, error) {
	return newsletterEntity.News{}, nil
}
func (r *inMemoryRepository) DeleteNews(newsID string) error {
	return nil
}
