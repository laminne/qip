package gormRepository

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/entity"
	"github.com/approvers/qip/pkg/utils/id"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) FindUsersByName(name string) ([]domain.User, error) {
	var u []entity.User
	res := r.db.Where(&entity.User{Name: name}).Find(&u)
	if res.Error != nil {
		return nil, res.Error
	}
	re := make([]domain.User, len(u))
	for i, v := range u {
		re[i] = *r.eToD(v)
	}
	return re, nil
}

func (r UserRepository) FindUserByID(id id.SnowFlakeID) (*domain.User, error) {
	u := entity.User{ID: string(id)}
	res := r.db.Debug().Take(&u)
	if res.Error != nil {
		return nil, res.Error
	}
	return r.eToD(u), nil
}

func (r UserRepository) FindUsersByInstanceID(id id.SnowFlakeID) ([]domain.User, error) {
	var u []entity.User
	res := r.db.Where("instanceid = ?", id).Find(&u)
	if res.Error != nil {
		return nil, res.Error
	}
	re := make([]domain.User, len(u))
	for i, v := range u {
		re[i] = *r.eToD(v)
	}
	return re, nil
}

func (r UserRepository) CreateUser(u domain.User) error {
	e := r.dToE(u)
	res := r.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r UserRepository) dToE(u domain.User) entity.User {
	return entity.User{
		ID:          string(u.GetID()),
		Name:        u.GetName(),
		DisplayName: u.GetDisplayName(),
		Role: func() int {
			if u.IsAdmin() {
				return 1
			}
			return 0
		}(),
		Bio:          u.GetBio(),
		IsFroze:      u.IsFroze(),
		InboxURL:     u.GetInboxURL(),
		OutboxURL:    u.GetOutboxURL(),
		FollowURL:    u.GetFollowURL(),
		FollowersURL: u.GetFollowersURL(),
		SecretKey:    u.GetSecretKey(),
		PublicKey:    u.GetPublicKey(),
		Password:     u.GetPassword(),
		IsLocalUser:  u.IsLocalUser(),
		CreatedAt:    u.GetCreatedAt(),
		UpdatedAt:    u.GetUpdatedAt(),
	}
}

func (r UserRepository) eToD(e entity.User) *domain.User {
	u, _ := domain.NewUser(id.SnowFlakeID(e.ID), e.Name, id.SnowFlakeID(e.InstanceID), e.IsLocalUser, e.CreatedAt)

	u.SetDisplayName(e.DisplayName)
	_, _ = u.SetRole(e.Role)
	if e.Bio != nil {
		u.SetBio(e.Bio)
	}
	if e.IsFroze {
		_, _ = u.Freeze()
	}
	_, _ = u.SetInboxURL(e.InboxURL)
	_, _ = u.SetOutboxURL(e.OutboxURL)
	_, _ = u.SetFollowerURL(e.FollowURL)
	_, _ = u.SetFollowerURL(e.FollowersURL)
	if e.IsLocalUser {
		_, _ = u.SetSecretKey(*e.SecretKey)
		_, _ = u.SetPassword(*e.Password)
	}
	_, _ = u.SetPublicKey(e.PublicKey)
	if e.IconImageID != nil {
		u.SetIcon(id.SnowFlakeID(*e.IconImageID))
	}
	if e.HeaderImageID != nil {
		u.SetHeader(id.SnowFlakeID(*e.HeaderImageID))
	}

	return u
}
