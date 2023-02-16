package domain

import (
	"errors"
	"time"

	"github.com/approvers/qip/pkg/utils/id"
)

type User struct {
	ID             id.SnowFlakeID
	Host           *string
	Name           string
	ScreenName     string
	Summary        string
	Password       string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	PrivateKey     string
	PublicKey      string
	WatcherCount   int
	WatchingCount  int
	PostsCount     int
	HeaderImageURL *string
	IconImageURL   *string
}

func NewUser(
	ID id.SnowFlakeID,
	host *string,
	name string,
	screenName string,
	summary string,
	password string,
	createdAt time.Time,
	updatedAt *time.Time,
	privateKey string,
	publicKey string,
	watcherCount int,
	watchingCount int,
	postsCount int,
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
		Password:       password,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		PrivateKey:     privateKey,
		PublicKey:      publicKey,
		WatcherCount:   watcherCount,
		WatchingCount:  watchingCount,
		PostsCount:     postsCount,
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

// UpdatePassword ユーザーのパスワードを変更
func (u *User) UpdatePassword(p string) {
	u.Password = p
}

// PostsCountUp ユーザーのノート数を増やす
func (u *User) PostsCountUp() {
	u.PostsCount++
}

// PostsCountDown ユーザーのノート数を減らす
func (u *User) PostsCountDown() {
	u.PostsCount--
}

// FollowerCountUp ユーザーの被フォロー数を増やす
func (u *User) FollowerCountUp() {
	u.WatcherCount++
}

// FollowerCountDown ユーザーのフォロワー数を減らす
func (u *User) FollowerCountDown() {
	u.WatcherCount--
}

// FollowingCountUp ユーザーのフォロー数を増やす
func (u *User) FollowingCountUp() {
	u.WatchingCount++
}

// FollowingCountDown ユーザーのフォロー数を減らす
func (u *User) FollowingCountDown() {
	u.WatchingCount--
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
