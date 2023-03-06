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

	arg := CreateUserCommand{
		Name:       "test",
		InstanceID: "123123123",
		IsLocal:    false,
	}
	err := createUserService.Handle(arg)
	assert.Equal(t, nil, err)
}
