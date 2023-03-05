package user

import (
	"errors"
	"time"

	"github.com/approvers/qip/pkg/domain/service"

	"github.com/approvers/qip/pkg/domain"

	"github.com/approvers/qip/pkg/utils/id"
)

type CreateUserCommand struct {
	Id         id.SnowFlakeID
	Name       string
	InstanceID id.SnowFlakeID
	IsLocal    bool
}

type ICreateUserService interface {
	Handle(c CreateUserCommand) error
}

type CreateUserService struct {
	userService service.UserService
}

func NewCreateUserService(userService service.UserService) *CreateUserService {
	return &CreateUserService{userService: userService}
}

func (s *CreateUserService) Handle(c CreateUserCommand) error {
	// ToDo: ここでCreatedAtを作っていいのか？
	now := time.Now()

	u, _ := domain.NewUser(c.Id, c.Name, c.InstanceID, c.IsLocal, now)
	// ToDo: 重複確認する
	if s.userService.Exists(u) {
		return errors.New("")
	}

	return nil
}
