package dummy

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/errorType"
	"github.com/laminne/qip/pkg/utils/id"
)

type InstanceRepository struct {
	data []domain.Instance
}

func NewInstanceRepository(data []domain.Instance) *InstanceRepository {
	return &InstanceRepository{data: data}
}

func (r *InstanceRepository) CreateInstance(i domain.Instance) error {
	for _, v := range r.data {
		if i.GetID() == v.GetID() {
			return errorType.NewErrExists("DummyInstanceRepository", "Instance exists")
		}
	}

	r.data = append(r.data, i)

	return nil
}

func (r *InstanceRepository) FindByID(id id.SnowFlakeID) (*domain.Instance, error) {
	for _, v := range r.data {
		if id == v.GetID() {
			return &v, nil
		}
	}

	return nil, errorType.NewErrNotFound("DummyInstanceRepository", "Instance Not Found")
}

func (r *InstanceRepository) FindByHost(host string) (*domain.Instance, error) {
	for _, v := range r.data {
		if host == v.GetHost() {
			return &v, nil
		}
	}
	return nil, errorType.NewErrNotFound("DummyInstanceRepository", "Instance Not Found")
}
