package router

import (
	"time"

	"github.com/approvers/qip/pkg/domain"
)

var UserMockData []domain.User

func init() {
	a, _ := domain.NewUser("123", "test1", "456", false, time.Now())
	b, _ := domain.NewUser("777", "test1", "789", false, time.Now())
	c, _ := domain.NewUser("888", "test3", "012", true, time.Now())

	UserMockData = append(UserMockData, *a)
	UserMockData = append(UserMockData, *b)
	UserMockData = append(UserMockData, *c)
}
