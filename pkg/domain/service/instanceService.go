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
	r, _ := s.repository.FindByID(i.GetID())
	r2, _ := s.repository.FindByHost(i.GetHost())
	if r == nil && r2 == nil {
		return false
	}

	return true
}
