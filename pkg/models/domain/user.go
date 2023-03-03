package domain

import (
	"errors"
	"unicode/utf8"

	"github.com/approvers/qip/pkg/utils/id"
)

const (
	NormalUserRole = iota
	AdminUserRole
)

type UserRole = int

type User struct {
	ID            id.SnowFlakeID
	Name          string
	DisplayName   string
	Role          UserRole
	InstanceID    id.SnowFlakeID
	Bio           *string
	HeaderImageID *id.SnowFlakeID
	IconImageID   *id.SnowFlakeID
	IsFroze       bool
	InboxURL      string
	OutboxURL     string
	FollowURL     string
	FollowersURL  string
	SecretKey     *string
	PublicKey     string
	Password      *string
	IsLocalUser   bool
}

func NewUser(id id.SnowFlakeID, name string, instanceID id.SnowFlakeID, isLocalUser bool) (*User, error) {
	if utf8.RuneCountInString(name) > 64 || utf8.RuneCountInString(name) < 0 {
		return nil, errors.New("ユーザー名の長さが制限を超えています")
	}

	return &User{
		ID:          id,
		Name:        name,
		InstanceID:  instanceID,
		Role:        NormalUserRole, // デフォルトは一般ユーザー
		IsLocalUser: isLocalUser,
		IsFroze:     false,
	}, nil
}

// SetDisplayName ユーザーの表示名を設定
func (u *User) SetDisplayName(displayName string) *User {
	if utf8.RuneCountInString(displayName) > 64 {
		// 切り捨てる
		u.DisplayName = displayName[:64]
	}

	if utf8.RuneCountInString(displayName) == 0 {
		// 指定がない場合はユーザー名にフォールバック
		u.DisplayName = u.Name
	}

	return u
}

// SetRole ユーザーのロールを設定
func (u *User) SetRole(role UserRole) (*User, error) {
	// ローカルユーザーである場合のみ設定できる
	if !u.IsLocalUser {
		return nil, errors.New("リモートユーザーのロールは変更できません")
	}

	u.Role = role
	return u, nil
}

// SetBio ユーザーの自己紹介文を設定
func (u *User) SetBio(bio *string) *User {
	if utf8.RuneCountInString(*bio) > 2000 {
		b := (*bio)[:2000]
		u.Bio = &b
		return u
	}

	u.Bio = bio
	return u
}

// SetHeader ユーザーのヘッダー画像を設定
func (u *User) SetHeader(id id.SnowFlakeID) *User {
	u.HeaderImageID = &id
	return u
}

// SetIcon ユーザーのアイコン画像を設定
func (u *User) SetIcon(id id.SnowFlakeID) *User {
	u.IconImageID = &id
	return u
}

// Freeze ユーザーを凍結
func (u *User) Freeze() (*User, error) {
	if u.IsFroze {
		return nil, errors.New("すでにユーザーは凍結されています")
	}

	u.IsFroze = true
	return u, nil
}

// Unfreeze ユーザーを解凍(凍結解除)
func (u *User) Unfreeze() (*User, error) {
	if !u.IsFroze {
		return nil, errors.New("ユーザーは凍結されていません")
	}

	u.IsFroze = false
	return u, nil
}

// SetInboxURL ユーザーのInboxURLを設定
func (u *User) SetInboxURL(url string) (*User, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}

	u.InboxURL = url
	return u, nil
}

// SetOutboxURL ユーザーのOutboxURLを設定
func (u *User) SetOutboxURL(url string) (*User, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}

	u.OutboxURL = url
	return u, nil
}

// SetFollowURL ユーザーのフォローURLを設定
func (u *User) SetFollowURL(url string) (*User, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}

	u.FollowURL = url
	return u, nil
}

// SetFollowerURL ユーザーのフォロワーURLを設定
func (u *User) SetFollowerURL(url string) (*User, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}

	u.FollowersURL = url
	return u, nil
}

// SetSecretKey ユーザーの秘密鍵を設定
func (u *User) SetSecretKey(key string) (*User, error) {
	// ローカルユーザーにしか設定できない
	if !u.IsLocalUser {
		return nil, errors.New("リモートユーザーに秘密鍵は設定できません")
	}

	if len(key) <= 0 {
		return nil, errors.New("鍵が短すぎます")
	}

	u.SecretKey = &key
	return u, nil
}

// SetPublicKey ユーザーの公開鍵を設定
func (u *User) SetPublicKey(key string) (*User, error) {
	if len(key) <= 0 {
		return nil, errors.New("鍵が短すぎます")
	}

	u.PublicKey = key
	return u, nil
}

// SetPassword ユーザーのパスワードを設定
func (u *User) SetPassword(pass string) (*User, error) {
	// ローカルユーザーにしか設定できない
	if !u.IsLocalUser {
		return nil, errors.New("リモートユーザーにパスワードは設定できません")
	}

	if len(pass) <= 0 {
		return nil, errors.New("パスワードの文字数が短すぎます")
	}

	u.Password = &pass
	return u, nil
}
