package repository

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type PostRepository interface {
	Create(p domain.Post) error
	FindByID(id id.SnowFlakeID) (*domain.Post, error)
	FindByAuthorID(id id.SnowFlakeID) ([]domain.Post, error)
	FindByIDWithUserIcon(id id.SnowFlakeID) (*PostUserFileJoinedData, error)
	FindByAuthorIDWithUserIcon(id id.SnowFlakeID) ([]PostUserFileJoinedData, error)
}

type PostUserFileJoinedData struct {
	Post *domain.Post
	User *domain.User
	File *domain.File
}
