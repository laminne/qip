package repository

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type FileRepository interface {
	Create(f domain.File) error
	FindByID(id id.SnowFlakeID) (*domain.File, error)
	FindByUploaderID(id id.SnowFlakeID) ([]domain.File, error)
}
