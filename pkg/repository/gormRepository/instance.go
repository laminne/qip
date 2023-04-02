package gormRepository

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/entity"
	"github.com/laminne/qip/pkg/utils/id"
	"gorm.io/gorm"
)

type InstanceRepository struct {
	db *gorm.DB
}

func NewInstanceRepository(db *gorm.DB) *InstanceRepository {
	return &InstanceRepository{db: db}
}

func (r *InstanceRepository) CreateInstance(i domain.Instance) error {
	e := r.dToE(i)
	res := r.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *InstanceRepository) FindByID(i id.SnowFlakeID) (*domain.Instance, error) {
	e := entity.Instance{ID: string(i)}
	res := r.db.Take(&e)
	if res.Error != nil {
		return &domain.Instance{}, res.Error
	}
	re := r.eToD(e)
	return &re, nil
}

func (r *InstanceRepository) FindByHost(host string) (*domain.Instance, error) {
	i := entity.Instance{Host: host}
	res := r.db.Take(&i)
	if res.Error != nil {
		return &domain.Instance{}, res.Error
	}
	re := r.eToD(i)
	return &re, nil
}

func (r *InstanceRepository) eToD(e entity.Instance) domain.Instance {
	re, _ := domain.NewInstance(id.SnowFlakeID(e.ID), e.Host, e.CreatedAt)
	_, _ = re.SetName(e.Name)
	_, _ = re.SetSoftwareName(e.SoftwareName)
	_, _ = re.SetSoftwareVersion(e.SoftwareVersion)
	_, _ = re.SetDescription(e.Description)
	if e.State == 1 {
		_, _ = re.InstanceBlock()
	}
	if e.UpdatedAt != nil {
		_, _ = re.SetUpdatedAt(*e.UpdatedAt)
	}
	return *re
}

func (r *InstanceRepository) dToE(d domain.Instance) entity.Instance {
	return entity.Instance{
		ID:              string(d.GetID()),
		Name:            d.GetName(),
		SoftwareName:    d.GetSoftwareName(),
		SoftwareVersion: d.GetSoftwareVersion(),
		Host:            d.GetHost(),
		Description:     d.GetDescription(),
		State: func() int {
			if d.IsInstanceBlocked() {
				return 1
			} else {
				return 0
			}
		}(),
		CreatedAt: d.GetCreatedAt(),
		UpdatedAt: d.GetUpdatedAt(),
	}
}
