package service

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/repository"
)

type PostService struct {
	repository repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repository: repo}
}

func (s *PostService) Exists(p *domain.Post) bool {
	res, err := s.repository.FindByID(p.GetID())
	if err != nil || res == nil {
		return false
	}

	return true
}
