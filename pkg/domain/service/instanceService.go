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
	res, err := s.repository.FindByID(i.GetID())
	if err != nil || res == nil {
		return false
	}

	return true
}
