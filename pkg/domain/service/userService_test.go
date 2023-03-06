package service

import (
	"testing"
	"time"

	"github.com/approvers/qip/pkg/repository/dummy"
	"github.com/stretchr/testify/assert"

	"github.com/approvers/qip/pkg/domain"
)

func TestUserService_Exists(t *testing.T) {
	testData, _ := domain.NewUser("112233", "test", "332211", false, time.Now())
	testDataArray := make([]domain.User, 0)
	testDataArray = append(testDataArray, *testData)
	r := dummy.NewUserRepository(testDataArray)

	userService := NewUserService(r)

	// 存在しないとき
	notExist, _ := domain.NewUser("445566", "hello", "332211", false, time.Now())
	assert.Equal(t, false, userService.Exists(notExist))

	// 存在するとき
	exist := testData
	assert.Equal(t, true, userService.Exists(exist))
}
