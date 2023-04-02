package post

import (
	"time"

	"github.com/laminne/qip/pkg/errorType"

	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/domain/service"
	"github.com/laminne/qip/pkg/repository"
	"github.com/laminne/qip/pkg/utils/id"
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
}

func NewCreatePostService(postService service.PostService, repository repository.PostRepository) *CreatePostService {
	idGenerator := id.NewSnowFlakeIDGenerator()
	return &CreatePostService{postService: postService, postRepository: repository, idGenerator: idGenerator}
}

func (s *CreatePostService) Handle(c CreatePostCommand) (*PostData, error) {
	now := time.Now()

	pID := s.idGenerator.NewID(now)
	p := domain.NewPost(pID, c.Body, c.Visibility, c.AuthorID, now)

	if s.postService.Exists(p) {
		return nil, errorType.NewErrExists("CreatePostService", "post exists")
	}

	err := s.postRepository.Create(*p)
	if err != nil {
		return nil, err
	}

	return NewPostData(*p), nil
}
