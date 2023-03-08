package repository

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type PostRepository interface {
	Create(p domain.Post) error
	FindByID(id id.SnowFlakeID) (*domain.Post, error)
	FindByAuthorID(id id.SnowFlakeID) ([]domain.Post, error)
}
