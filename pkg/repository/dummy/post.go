package dummy

import (
	"errors"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type PostRepository struct {
	data []domain.Post
}

func NewPostRepository(data []domain.Post) *PostRepository {
	return &PostRepository{data: data}
}

func (p *PostRepository) FindByID(id id.SnowFlakeID) (*domain.Post, error) {
	for _, v := range p.data {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("PostNotFound")
}

func (p *PostRepository) Create(post domain.Post) error {
	for _, v := range p.data {
		if v.GetID() == post.GetID() {
			return errors.New("PostExists")
		}
	}

	p.data = append(p.data, post)
	return nil
}

func (p *PostRepository) FindByAuthorID(id id.SnowFlakeID) ([]domain.Post, error) {
	res := make([]domain.Post, 0)
	for _, v := range p.data {
		if v.GetAuthorID() == id {
			res = append(res, v)
		}
	}
	return res, nil
}
