package bun

import (
	"context"
	"errors"
	"fmt"

	"github.com/laminne/notepod/pkg/models/domain"
	"github.com/laminne/notepod/pkg/models/entity"
	"github.com/laminne/notepod/pkg/utils/id"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) dtoe(u domain.User) entity.User {
	return entity.User{
		ID:             string(u.ID),
		Host:           u.Host,
		Name:           u.Name,
		ScreenName:     u.ScreenName,
		Summary:        u.Summary,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
		PrivateKey:     u.PrivateKey,
		PublicKey:      u.PublicKey,
		FollowerCount:  u.FollowerCount,
		FollowingCount: u.FollowingCount,
		NoteCount:      u.NoteCount,
		HeaderImageURL: u.HeaderImageURL,
		IconImageURL:   u.IconImageURL,
	}
}

func (r UserRepository) etod(e entity.User) domain.User {
	return domain.User{
		ID:             id.SnowFlakeID(e.ID),
		Host:           e.Host,
		Name:           e.Name,
		ScreenName:     e.ScreenName,
		Summary:        e.Summary,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
		PrivateKey:     e.PrivateKey,
		PublicKey:      e.PublicKey,
		FollowerCount:  e.FollowerCount,
		FollowingCount: e.FollowingCount,
		NoteCount:      e.NoteCount,
		HeaderImageURL: e.HeaderImageURL,
		IconImageURL:   e.IconImageURL,
	}
}

func (r UserRepository) Create(u domain.User) (domain.User, error) {
	e := r.dtoe(u)

	_, err := r.db.NewInsert().Model(&e).Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		return domain.User{}, errors.New("failed to save user")
	}

	return r.etod(e), nil
}

func (r UserRepository) Update(u domain.User) (domain.User, error) {
	e := r.dtoe(u)

	_, err := r.db.NewUpdate().Model(&e).Exec(context.Background())
	if err != nil {
		return domain.User{}, errors.New("failed to update user")
	}

	return r.etod(e), nil
}

func (r UserRepository) FindByID(id id.SnowFlakeID) (*domain.User, error) {
	u := entity.User{}
	err := r.db.NewSelect().Model(&u).Where("id = ?", id).Scan(context.Background(), &u)
	if err != nil {
		fmt.Println(err, "err")
		return nil, errors.New("failed to find user by id")
	}

	res := r.etod(u)
	return &res, nil
}

func (r UserRepository) FindByUserName(n string) ([]domain.User, error) {
	var u, b []entity.User

	err := r.db.NewSelect().Model(&u).Where(fmt.Sprintf("userName = %s", n)).Scan(context.Background(), &b)
	if err != nil {
		return nil, errors.New("failed to find user by username")
	}

	var res []domain.User
	for _, v := range b {
		res = append(res, r.etod(v))
	}
	return res, nil
}

func (r UserRepository) FindByHost(h string) ([]domain.User, error) {
	var u, b []entity.User

	err := r.db.NewSelect().Model(&u).Where(fmt.Sprintf("host = %s", h)).Scan(context.Background(), &b)
	if err != nil {
		return nil, errors.New("failed to find user by host")
	}

	var res []domain.User
	for _, v := range b {
		res = append(res, r.etod(v))
	}

	return res, nil
}

func (r UserRepository) FindLocalByUserName(n string) (domain.User, error) {
	var u, b []entity.User

	err := r.db.NewSelect().Model(&u).Where(fmt.Sprintf("host IS null")).Scan(context.Background(), &b)
	if err != nil {
		return domain.User{}, errors.New("failed to find user by username")
	}

	for _, v := range b {
		if v.Name == n {
			return r.etod(v), nil
		}
	}

	return domain.User{}, errors.New("user not found")
}
