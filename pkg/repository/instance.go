package repository

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type InstanceRepository interface {
	CreateInstance(i domain.Instance) error
	FindByID(i id.SnowFlakeID) (*domain.Instance, error)
	FindByHost(host string) (*domain.Instance, error)
}
