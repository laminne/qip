package instance

import (
	"github.com/approvers/qip/pkg/errorType"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type IFindInstanceService interface {
	FindByID(id id.SnowFlakeID) (*InstanceData, error)
}

type FindInstanceService struct {
	repository repository.InstanceRepository
}

func NewFindInstanceService(repo repository.InstanceRepository) *FindInstanceService {
	return &FindInstanceService{repository: repo}
}

func (s FindInstanceService) FindByID(id id.SnowFlakeID) (*InstanceData, error) {
	i, err := s.repository.FindByID(id)
	if err != nil {
		return nil, errorType.NewErrNotFound("FindInstanceService", "no such instance")
	}

	return NewInstanceData(*i), nil
}
