package gormRepository

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/entity"
	"github.com/approvers/qip/pkg/utils/id"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(d *gorm.DB) *PostRepository {
	return &PostRepository{d}
}

func (r *PostRepository) Create(p domain.Post) error {
	e := r.dToE(p)
	res := r.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *PostRepository) FindByID(id id.SnowFlakeID) (*domain.Post, error) {
	p := entity.Post{ID: string(id)}
	res := r.db.Take(&p)
	if res.Error != nil {
		return &domain.Post{}, res.Error
	}
	re := r.eToD(p)
	return &re, nil
}

func (r *PostRepository) FindByAuthorID(id id.SnowFlakeID) ([]domain.Post, error) {
	var p []entity.Post
	res := r.db.Where("authorid = ?", id).Find(&p)
	if res.Error != nil {
		return nil, res.Error
	}

	re := make([]domain.Post, len(p))
	for i, v := range p {
		re[i] = r.eToD(v)
	}
	return re, nil
}

func (r *PostRepository) dToE(p domain.Post) entity.Post {
	return entity.Post{
		ID:         string(p.GetID()),
		Body:       p.GetBody(),
		Visibility: p.GetVisibility(),
		AuthorID:   string(p.GetAuthorID()),
		CreatedAt:  p.GetCreatedAt(),
	}
}

func (r *PostRepository) eToD(e entity.Post) domain.Post {
	// ToDo: AttachmentFileを考慮する
	return *domain.NewPost(id.SnowFlakeID(e.ID), e.Body, e.Visibility, id.SnowFlakeID(e.AuthorID), e.CreatedAt)
}
