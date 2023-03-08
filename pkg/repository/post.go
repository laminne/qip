package repository

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type PostRepository interface {
	FindByID(id id.SnowFlakeID) (*domain.Post, error)
}
