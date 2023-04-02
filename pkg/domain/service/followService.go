package service

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/repository"
)

type FollowService struct {
	repository repository.UserRepository
}

func NewFollowService(repo repository.UserRepository) *FollowService {
	return &FollowService{repository: repo}
}

func (s *FollowService) Exists(f domain.Follow) bool {
	res, err := s.repository.FindUserFollow(f.GetUserID())
	if err != nil {
		return false
	}

	for _, v := range res {
		if v.GetTargetID() == f.GetTargetID() {
			return true
		}
	}

	return false
}
