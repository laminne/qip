package controller

import (
	"github.com/approvers/qip/pkg/application/post"
	"github.com/approvers/qip/pkg/application/user"
	"github.com/approvers/qip/pkg/controller/models"
	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type PostController struct {
	repo              repository.PostRepository
	userRepository    repository.UserRepository
	createPostService post.ICreatePostService
	findPostService   post.IFindPostService
	findUserService   user.IFindUserService
}

func NewPostController(r repository.PostRepository, rp repository.UserRepository) *PostController {
	ps := post.NewCreatePostService(*service.NewPostService(r), r)
	return &PostController{
		repo:              r,
		createPostService: ps,
		findPostService:   post.NewFindPostService(r),
		findUserService:   user.NewFindUserService(rp),
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
	res, err := p.findPostService.FindByIDWithUserIcon(id.SnowFlakeID(pID))
	if err != nil {
		return models.GetPostResponseJSON{}, err
	}

	return models.GetPostResponseJSON{
		ID:         string(res.PostData.GetID()),
		Body:       res.PostData.GetBody(),
		AuthorID:   string(res.PostData.GetAuthorID()),
		Visibility: p.visibilityConverter(res.PostData.GetVisibility()),
		CreatedAt:  res.PostData.GetCreatedAt(),
		User: models.GetPostResponseAuthorData{
			Name:       res.UserData.Name(),
			ScreenName: res.UserData.DisplayName(),
			IconURL:    res.FileData.FileURL(),
		},
	}, nil
}

func (p *PostController) FindByAuthorID(uID string) ([]models.GetPostResponseJSON, error) {
	res, err := p.findPostService.FindByAuthorIDWithUserIcon(id.SnowFlakeID(uID))
	if err != nil {
		return []models.GetPostResponseJSON{}, err
	}

	resp := make([]models.GetPostResponseJSON, len(res))

	for i, v := range res {
		resp[i] = models.GetPostResponseJSON{
			ID:         string(v.PostData.GetID()),
			Body:       v.PostData.GetBody(),
			AuthorID:   string(v.PostData.GetAuthorID()),
			Visibility: p.visibilityConverter(v.GetVisibility()),
			CreatedAt:  v.PostData.GetCreatedAt(),
			User: models.GetPostResponseAuthorData{
				Name:       v.UserData.Name(),
				ScreenName: v.UserData.DisplayName(),
				IconURL:    v.FileData.FileURL(),
			},
		}
	}

	return resp, nil
}
