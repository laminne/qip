package service

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (s *UserService) Exists(u *domain.User) bool {
	/*
		仕様:
		- 同じユーザー名は各インスタンスで1人しか存在できない
	*/
	users, _ := s.repository.FindUsersByName(u.GetName())
	for _, v := range users {
		if v.GetInstanceID() == u.GetInstanceID() {
			return true
		}
	}
	return false
}
