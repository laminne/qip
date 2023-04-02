package user

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/laminne/qip/pkg/repository"
	"github.com/laminne/qip/pkg/repository/dummy"

	"github.com/laminne/qip/pkg/domain"
)

var (
	data            []domain.User
	repo            repository.UserRepository
	findUserService FindUserService
)

func init() {
	fmt.Println("Init")
	a, _ := domain.NewUser("123", "test1", "456", false, time.Now())
	b, _ := domain.NewUser("777", "test1", "789", false, time.Now())
	c, _ := domain.NewUser("888", "test3", "012", true, time.Now())
	data = append(data, *a)
	data = append(data, *b)
	data = append(data, *c)

	repo := dummy.NewUserRepository(data)
	findUserService = *NewFindUserService(repo)
}

func TestFindUserService_FindByID(t *testing.T) {
	// 成功するとき
	a, _ := findUserService.FindByID("123")
	assert.Equal(t, NewUserData(data[0]), a)

	// 失敗するとき
	_, err := findUserService.FindByID("000")
	assert.NotEqual(t, nil, err)
}

func TestFindUserService_FindByInstanceID(t *testing.T) {
	ex := []UserData{*NewUserData(data[1])}
	// 成功するとき
	a, _ := findUserService.FindByInstanceID("789")
	assert.Equal(t, ex, a)
	// 失敗するとき
	b, _ := findUserService.FindByInstanceID("777")
	assert.Equal(t, 0, len(b))
}

func TestFindUserService_FindByName(t *testing.T) {
	ex := []UserData{*NewUserData(data[0]), *NewUserData(data[1])}

	// 成功するとき
	a, _ := findUserService.FindByName("test1")
	assert.Equal(t, ex, a)

	// 失敗するとき
	b, _ := findUserService.FindByName("test0")
	assert.Equal(t, 0, len(b))
}
