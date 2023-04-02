package instance

import (
	"time"

	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/domain/service"
	"github.com/laminne/qip/pkg/errorType"
	"github.com/laminne/qip/pkg/repository"
	"github.com/laminne/qip/pkg/utils/id"
)

type CreateInstanceCommand struct {
	Host string
}

type ICreateInstanceService interface {
	Handle(c CreateInstanceCommand) (*InstanceData, error)
}

type CreateInstanceService struct {
	instanceService    service.InstanceService
	instanceRepository repository.InstanceRepository
	idGenerator        id.Generator
}

func NewCreateInstanceService(s service.InstanceService, r repository.InstanceRepository) *CreateInstanceService {
	idGenerator := id.NewSnowFlakeIDGenerator()
	return &CreateInstanceService{
		instanceService:    s,
		instanceRepository: r,
		idGenerator:        idGenerator,
	}
}

func (s *CreateInstanceService) Handle(c CreateInstanceCommand) (*InstanceData, error) {
	now := time.Now()

	iID := s.idGenerator.NewID(now)
	i, _ := domain.NewInstance(iID, c.Host, now)
	if s.instanceService.Exists(*i) {
		return nil, errorType.NewErrExists("CreateInstanceService", "instance exists")
	}

	err := s.instanceRepository.CreateInstance(*i)
	if err != nil {
		return nil, err
	}

	return NewInstanceData(*i), nil
}
