package post

import (
	"github.com/laminne/qip/pkg/application/file"
	"github.com/laminne/qip/pkg/application/user"
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/errorType"
	"github.com/laminne/qip/pkg/repository"
	"github.com/laminne/qip/pkg/utils/id"
)

type IFindPostService interface {
	FindByID(id id.SnowFlakeID) (*PostData, error)
	FindByAuthorID(id id.SnowFlakeID) ([]PostData, error)
	FindByIDWithUserIcon(id id.SnowFlakeID) (*PostWithUserData, error)
	FindByAuthorIDWithUserIcon(id id.SnowFlakeID) ([]PostWithUserData, error)
}

type PostWithUserData struct {
	PostData
	user.UserData
	file.FileData
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

func (f *FindPostService) FindByIDWithUserIcon(id id.SnowFlakeID) (*PostWithUserData, error) {
	r, err := f.postRepository.FindByIDWithUserIcon(id)
	if err != nil {
		return nil, errorType.NewErrNotFound("FindPostService", "post not found")
	}

	if r == nil || r.Post == nil {
		return nil, errorType.NewErrNotFound("FindPostService", "post not found")
	}

	res := &PostWithUserData{
		PostData: *NewPostData(*r.Post),
		UserData: *user.NewUserData(*r.User),
		FileData: *file.NewFileData(*r.File),
	}
	return res, nil
}

func (f *FindPostService) FindByAuthorIDWithUserIcon(id id.SnowFlakeID) ([]PostWithUserData, error) {
	r, err := f.postRepository.FindByAuthorIDWithUserIcon(id)
	if err != nil {
		return nil, errorType.NewErrNotFound("FindPostService", "post not found")
	}

	res := make([]PostWithUserData, len(r))
	for j, k := range r {
		res[j] = PostWithUserData{
			PostData: *NewPostData(*k.Post),
			UserData: *user.NewUserData(*k.User),
			FileData: *file.NewFileData(*k.File),
		}
	}

	return res, nil
}
