package post

import (
	"fmt"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
	"github.com/approvers/qip/pkg/utils/logger"
)

type IFindPostService interface {
	FindByID(id id.SnowFlakeID) (*PostData, error)
	FindByAuthorID(id id.SnowFlakeID) ([]PostData, error)
}

type FindPostService struct {
	postRepository repository.PostRepository
	logger         logger.Logger
}

func NewFindPostService(repo repository.PostRepository, log logger.Logger) *FindPostService {
	return &FindPostService{postRepository: repo, logger: log}
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
		if repository.NotFound == err {
			f.logger.Error(fmt.Sprintf("[FindPostService] No such post (PostID: %v)", id))
		}
		return nil, fmt.Errorf("post not found, %w", err)
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
