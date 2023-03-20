package gormRepository

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/entity"
	"github.com/approvers/qip/pkg/utils/id"
	"gorm.io/gorm"
)

type FollowRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) *FollowRepository {
	return &FollowRepository{db: db}
}

func (r *FollowRepository) Create(f domain.Follow) error {
	e := r.dToE(f)
	res := r.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *FollowRepository) FindUserFollowers(i id.SnowFlakeID) ([]domain.Follow, error) {
	e := entity.Follow{TargetID: string(i)}
	var f []entity.Follow
	res := r.db.Where(&e).Find(&f)
	if res.Error != nil {
		return nil, res.Error
	}
	re := make([]domain.Follow, len(f))
	for i, v := range f {
		re[i] = r.eToD(v)
	}
	return re, nil
}

func (r *FollowRepository) FindUserFollow(i id.SnowFlakeID) ([]domain.Follow, error) {
	e := entity.Follow{UserID: string(i)}
	var f []entity.Follow
	res := r.db.Where(&e).Find(&f)
	if res.Error != nil {
		return nil, res.Error
	}
	re := make([]domain.Follow, len(f))
	for i, v := range f {
		re[i] = r.eToD(v)
	}
	return re, nil
}

func (r *FollowRepository) dToE(d domain.Follow) entity.Follow {
	return entity.Follow{
		UserID:    string(d.GetUserID()),
		TargetID:  string(d.GetTargetID()),
		CreatedAt: d.GetCreatedAt(),
	}
}

func (r *FollowRepository) eToD(e entity.Follow) domain.Follow {
	d, _ := domain.NewFollow(id.SnowFlakeID(e.UserID), id.SnowFlakeID(e.TargetID), e.CreatedAt)
	return *d
}

func (r *FollowRepository) UnFollow(from id.SnowFlakeID, target id.SnowFlakeID) error {
	res := r.db.Where(&entity.Follow{UserID: string(from), TargetID: string(target)}).Delete(&entity.Follow{UserID: string(from), TargetID: string(target)})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
