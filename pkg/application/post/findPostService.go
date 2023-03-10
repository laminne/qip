package post

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/errorType"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type IFindPostService interface {
	FindByID(id id.SnowFlakeID) (*PostData, error)
	FindByAuthorID(id id.SnowFlakeID) ([]PostData, error)
}

type FindPostService struct {
	postRepository repository.PostRepository
}

func NewFindPostService(repo repository.PostRepository) *FindPostService {
	return &FindPostService{postRepository: repo}
}

func (f *FindPostService) convert(p []domain.Post) []PostData {
	d := make([]PostData, len(p))

	for i, v := range p {
		d[i] = *NewPostData(v)
	}
	return d
}

func (f *FindPostService) FindByID(id id.SnowFlakeID) (*PostData, error) {
	p, err := f.postRepository.FindByID(id)
	if err != nil {
		return nil, errorType.NewErrNotFound("FindPostService", "post not found")
	}

	return NewPostData(*p), nil
}

func (f *FindPostService) FindByAuthorID(id id.SnowFlakeID) ([]PostData, error) {
	p, err := f.postRepository.FindByAuthorID(id)
	if err != nil {
		return nil, err
	}

	return f.convert(p), nil
}
