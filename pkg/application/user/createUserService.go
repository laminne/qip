package user

import (
	"errors"
	"time"

	"github.com/laminne/qip/pkg/errorType"

	"github.com/laminne/qip/pkg/utils/key"

	"github.com/laminne/qip/pkg/utils/password/argon2"

	"github.com/laminne/qip/pkg/utils/password"

	"github.com/laminne/qip/pkg/repository"

	"github.com/laminne/qip/pkg/domain/service"

	"github.com/laminne/qip/pkg/domain"

	"github.com/laminne/qip/pkg/utils/id"
)

type CreateUserCommand struct {
	Name       string
	InstanceID id.SnowFlakeID
	IsLocal    bool
	Password   string
}

type ICreateUserService interface {
	Handle(c CreateUserCommand) error
}

type CreateUserService struct {
	userService     service.UserService
	userRepository  repository.UserRepository
	idGenerator     id.Generator
	passwordEncoder password.Encoder
}

func NewCreateUserService(userService service.UserService, repository repository.UserRepository) *CreateUserService {
	idGenerator := id.NewSnowFlakeIDGenerator()
	encoder := argon2.NewArgon2PasswordEncoder()
	return &CreateUserService{userService: userService, idGenerator: idGenerator, userRepository: repository, passwordEncoder: encoder}
}

func (s *CreateUserService) Handle(c CreateUserCommand) error {
	// ToDo: ここでCreatedAtを作っていいのか？
	now := time.Now()

	uID := s.idGenerator.NewID(now)
	u, _ := domain.NewUser(uID, c.Name, c.InstanceID, c.IsLocal, now)

	if s.userService.Exists(u) {
		return errors.New("")
	}

	if u.IsLocalUser() {
		if len(c.Password) == 0 {
			return errorType.NewErrMissingPassword("CreateUserService", "password required")
		}

		pw := c.Password
		encoded, _ := s.passwordEncoder.EncodePassword(pw)
		_, err := u.SetPassword(string(encoded))
		if err != nil {
			return err
		}

		keys, err := key.GenRSAKey()
		if err != nil {
			return err
		}

		_, err = u.SetPublicKey(string(keys.PublicKey))
		if err != nil {
			return err
		}

		_, err = u.SetSecretKey(string(keys.PrivateKey))
		if err != nil {
			return err
		}
	}

	err := s.userRepository.CreateUser(*u)
	if err != nil {
		return err
	}

	return nil
}
