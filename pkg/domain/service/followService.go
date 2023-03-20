package service

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/repository"
)

type FollowService struct {
	repository repository.FollowRepository
}

func NewFollowService(repo repository.FollowRepository) *FollowService {
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
