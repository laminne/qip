package user

import (
	"testing"

	"github.com/approvers/qip/pkg/domain"

	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository/dummy"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserService_Handle(t *testing.T) {
	d := new([]domain.User)
	repository := dummy.NewUserRepository(*d)
	userService := service.NewUserService(repository)
	createUserService := NewCreateUserService(*userService, repository)

	// 成功するとき
	arg := CreateUserCommand{
		Name:       "test",
		InstanceID: "123123123",
		IsLocal:    true,
		Password:   "password@123",
	}
	err := createUserService.Handle(arg)
	assert.Equal(t, nil, err)

	// 失敗するとき
	// ローカルユーザーでパスワードを設定しないことはできない
	arg2 := CreateUserCommand{
		Name:       "test3",
		InstanceID: "123123123",
		IsLocal:    true,
		Password:   "",
	}
	err2 := createUserService.Handle(arg2)
	assert.NotEqual(t, nil, err2)

}
