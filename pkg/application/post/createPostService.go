package post

import (
	"errors"
	"fmt"
	"time"

	"github.com/approvers/qip/pkg/utils/logger"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type CreatePostCommand struct {
	Body       string
	AuthorID   id.SnowFlakeID
	Visibility domain.PostVisibility
}

type ICreatePostService interface {
	Handle(c CreatePostCommand) (*PostData, error)
}

type CreatePostService struct {
	postService    service.PostService
	postRepository repository.PostRepository
	idGenerator    id.Generator
	logger         logger.Logger
}

func NewCreatePostService(postService service.PostService, repository repository.PostRepository, log logger.Logger) *CreatePostService {
	idGenerator := id.NewSnowFlakeIDGenerator()
	return &CreatePostService{postService: postService, postRepository: repository, idGenerator: idGenerator, logger: log}
}

func (s *CreatePostService) Handle(c CreatePostCommand) (*PostData, error) {
	now := time.Now()

	pID := s.idGenerator.NewID(now)
	p := domain.NewPost(pID, c.Body, c.Visibility, c.AuthorID, now)

	if s.postService.Exists(p) {
		s.logger.Error(fmt.Sprintf("[CreatePostService] PostID Duplicated ID: %v", p.GetID()))
		return nil, errors.New("IDが重複しています")
	}

	err := s.postRepository.Create(*p)
	if err != nil {
		return nil, err
	}

	return NewPostData(*p), nil
}
