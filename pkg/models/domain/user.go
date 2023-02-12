package models

import (
	"errors"
	"time"

	"github.com/laminne/notepod/pkg/utils/id"
)

type User struct {
	ID             id.SnowFlakeID
	Host           *string
	Name           string
	ScreenName     string
	Summary        string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	PrivateKey     string
	PublicKey      string
	FollowerCount  int
	FollowingCount int
	NoteCount      int
	HeaderImageURL *string
	IconImageURL   *string
}

func NewUser(
	ID id.SnowFlakeID,
	host *string,
	name string,
	screenName string,
	summary string,
	createdAt time.Time,
	updatedAt *time.Time,
	privateKey string,
	publicKey string,
	followerCount int,
	followingCount int,
	noteCount int,
	headerImageURL *string,
	iconImageURL *string) (*User, error) {

	if err := validateUserName(name); err != nil {
		return nil, errors.New("username validation failed")
	}

	return &User{
		ID:             ID,
		Host:           host,
		Name:           name,
		ScreenName:     screenName,
		Summary:        summary,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		PrivateKey:     privateKey,
		PublicKey:      publicKey,
		FollowerCount:  followerCount,
		FollowingCount: followingCount,
		NoteCount:      noteCount,
		HeaderImageURL: headerImageURL,
		IconImageURL:   iconImageURL,
	}, nil
}

func validateUserName(n string) error {
	if len(n) > 64 || len(n) <= 0 {
		return errors.New("username validation failed")
	}
	return nil
}

// NoteCountUp ユーザーのノート数を増やす
func (u *User) NoteCountUp() {
	u.NoteCount++
}

// NoteCountDown ユーザーのノート数を減らす
func (u *User) NoteCountDown() {
	u.NoteCount--
}

// FollowerCountUp ユーザーの被フォロー数を増やす
func (u *User) FollowerCountUp() {
	u.FollowerCount++
}

// FollowerCountDown ユーザーのフォロワー数を減らす
func (u *User) FollowerCountDown() {
	u.FollowerCount--
}

// FollowingCountUp ユーザーのフォロー数を増やす
func (u *User) FollowingCountUp() {
	u.FollowingCount++
}

// FollowingCountDown ユーザーのフォロー数を減らす
func (u *User) FollowingCountDown() {
	u.FollowingCount--
}

// UpdateUserSummary ユーザーの自己紹介を更新
func (u *User) UpdateUserSummary(s string) error {
	if len(s) > 256 {
		return errors.New("user summary too long")
	}

	u.Summary = s
	return nil
}

// UpdateUserScreenName ユーザーの表示名を更新
func (u *User) UpdateUserScreenName(s string) error {
	if len(s) > 32 || len(s) == 0 {
		return errors.New("username too short or too long")
	}

	u.ScreenName = s
	return nil
}

// UpdateUserHeaderImage ヘッダー画像を更新
func (u *User) UpdateUserHeaderImage(url string) {
	u.HeaderImageURL = &url
}

// UpdateUserIconImage アイコン画像を更新
func (u *User) UpdateUserIconImage(url string) {
	u.IconImageURL = &url
}
