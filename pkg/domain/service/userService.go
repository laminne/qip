package service

import "github.com/approvers/qip/pkg/domain"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Exists(u *domain.User) bool {
	/*
		仕様:
		- 同じユーザー名は各インスタンスで1人しか存在できない
	*/
	// ToDo: Repositoryからデータを取ってくる
	return true
}
