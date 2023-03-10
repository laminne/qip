package service

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/repository"
)

type InstanceService struct {
	repository repository.InstanceRepository
}

func NewInstanceService(repo repository.InstanceRepository) *InstanceService {
	return &InstanceService{repository: repo}
}

func (s *InstanceService) Exists(i domain.Instance) bool {
	r, err := s.repository.FindByID(i.GetID())
	r2, err := s.repository.FindByHost(i.GetHost())
	if err != nil {
		return false
	}
	if r == nil && r2 == nil {
		return false
	}

	return true
}
