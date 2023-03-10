package router

import (
	"time"

	"github.com/approvers/qip/pkg/domain"
)

var UserMockData []domain.User
var PostMockData []domain.Post

func init() {
	func() {
		a, _ := domain.NewUser("123", "test1", "456", false, time.Now())
		b, _ := domain.NewUser("777", "test1", "789", false, time.Now())
		c, _ := domain.NewUser("888", "test3", "012", true, time.Now())

		_, _ = c.SetPassword("Argon2.36861cacfd0197139119336a706ac7e2b8f49456.305[0]")

		UserMockData = append(UserMockData, *a)
		UserMockData = append(UserMockData, *b)
		UserMockData = append(UserMockData, *c)
	}()

	func() {
		a := domain.NewPost("111", "Hello world", domain.Home, "222", time.Now())
		b := domain.NewPost("333", "こんにちは世界", domain.Global, "555", time.Now())
		c := domain.NewPost("777", "おはようございます", domain.Follower, "999", time.Now())

		PostMockData = append(PostMockData, *a)
		PostMockData = append(PostMockData, *b)
		PostMockData = append(PostMockData, *c)
	}()
}
