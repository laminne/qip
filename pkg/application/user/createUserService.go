package user

import (
	"errors"
	"time"

	"github.com/approvers/qip/pkg/repository"

	"github.com/approvers/qip/pkg/domain/service"

	"github.com/approvers/qip/pkg/domain"

	"github.com/approvers/qip/pkg/utils/id"
)

type CreateUserCommand struct {
	Name       string
	InstanceID id.SnowFlakeID
	IsLocal    bool
}

type ICreateUserService interface {
	Handle(c CreateUserCommand) error
}

type CreateUserService struct {
	userService    service.UserService
	userRepository repository.UserRepository
	idGenerator    id.Generator
}

func NewCreateUserService(userService service.UserService, repository repository.UserRepository) *CreateUserService {
	idGenerator := id.NewSnowFlakeIDGenerator()
	return &CreateUserService{userService: userService, idGenerator: idGenerator, userRepository: repository}
}

func (s *CreateUserService) Handle(c CreateUserCommand) error {
	// ToDo: ここでCreatedAtを作っていいのか？
	now := time.Now()

	uID := s.idGenerator.NewID(now)
	u, _ := domain.NewUser(uID, c.Name, c.InstanceID, c.IsLocal, now)

	if s.userService.Exists(u) {
		return errors.New("")
	}

	err := s.userRepository.CreateUser(*u)
	if err != nil {
		return err
	}

	return nil
}
