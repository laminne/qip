package controller

import (
	"github.com/approvers/qip/pkg/application/post"
	"github.com/approvers/qip/pkg/controller/models"
	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type PostController struct {
	repo              repository.PostRepository
	createPostService post.ICreatePostService
	findPostService   post.IFindPostService
}

func NewPostController(r repository.PostRepository) *PostController {
	ps := post.NewCreatePostService(*service.NewPostService(r), r)
	return &PostController{
		repo:              r,
		createPostService: ps,
		findPostService:   post.NewFindPostService(r),
	}
}

func (p *PostController) Create(body string, authorID id.SnowFlakeID, visibility post.Visibility) (models.CreatePostResponseJSON, error) {
	c := post.CreatePostCommand{
		Body:       body,
		AuthorID:   authorID,
		Visibility: visibility,
	}

	res, err := p.createPostService.Handle(c)
	if err != nil {
		return models.CreatePostResponseJSON{}, err
	}

	return models.CreatePostResponseJSON{
		ID:         string(res.GetID()),
		Body:       res.GetBody(),
		AuthorID:   string(res.GetAuthorID()),
		Visibility: p.visibilityConverter(res.GetVisibility()),
		CreatedAt:  res.GetCreatedAt(),
	}, nil
}

func (p *PostController) visibilityConverter(v int) string {
	switch v {
	case post.Global:
		return "global"
	case post.Home:
		return "home"
	case post.Follower:
		return "follower"
	case post.Direct:
		return "direct"
	}
	return ""
}

func (p *PostController) FindByID(pID string) (models.GetPostResponseJSON, error) {
	res, err := p.findPostService.FindByID(id.SnowFlakeID(pID))
	if err != nil {
		return models.GetPostResponseJSON{}, err
	}

	return models.GetPostResponseJSON{
		ID:         string(res.GetID()),
		Body:       res.GetBody(),
		AuthorID:   string(res.GetAuthorID()),
		Visibility: p.visibilityConverter(res.GetVisibility()),
		CreatedAt:  res.GetCreatedAt(),
	}, nil
}
