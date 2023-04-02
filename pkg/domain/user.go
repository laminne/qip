package domain

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/laminne/qip/pkg/utils/id"
)

const (
	NormalUserRole = iota
	AdminUserRole
)

type UserRole = int

type User struct {
	id            id.SnowFlakeID
	name          string
	displayName   string
	role          UserRole
	instanceID    id.SnowFlakeID
	bio           *string
	headerImageID *id.SnowFlakeID
	iconImageID   *id.SnowFlakeID
	isFroze       bool
	inboxURL      string
	outboxURL     string
	followURL     string
	followersURL  string
	secretKey     *string
	publicKey     string
	password      *string
	isLocalUser   bool
	createdAt     time.Time
	updatedAt     *time.Time
}

func NewUser(id id.SnowFlakeID, name string, instanceID id.SnowFlakeID, isLocalUser bool, now time.Time) (*User, error) {
	if utf8.RuneCountInString(name) > 64 || utf8.RuneCountInString(name) < 0 {
		return nil, errors.New("ユーザー名の長さが制限を超えています")
	}

	return &User{
		id:          id,
		name:        name,
		instanceID:  instanceID,
		role:        NormalUserRole, // デフォルトは一般ユーザー
		isLocalUser: isLocalUser,
		isFroze:     false,
		createdAt:   now,
	}, nil
}

// SetDisplayName ユーザーの表示名を設定
func (u *User) SetDisplayName(displayName string) *User {
	if utf8.RuneCountInString(displayName) > 64 {
		// 切り捨てる
		u.displayName = displayName[:64]
	} else if utf8.RuneCountInString(displayName) == 0 {
		// 指定がない場合はユーザー名にフォールバック
		u.displayName = u.name
	} else {
		u.displayName = displayName
	}

	return u
}

// SetRole ユーザーのロールを設定
func (u *User) SetRole(role UserRole) (*User, error) {
	// ローカルユーザーである場合のみ設定できる
	if !u.isLocalUser {
		return nil, errors.New("リモートユーザーのロールは変更できません")
	}

	u.role = role
	return u, nil
}

// SetBio ユーザーの自己紹介文を設定
func (u *User) SetBio(bio *string) *User {
	if utf8.RuneCountInString(*bio) > 2000 {
		b := (*bio)[:2000]
		u.bio = &b
		return u
	}

	u.bio = bio
	return u
}

// SetHeader ユーザーのヘッダー画像を設定
func (u *User) SetHeader(id id.SnowFlakeID) *User {
	u.headerImageID = &id
	return u
}

// SetIcon ユーザーのアイコン画像を設定
func (u *User) SetIcon(id id.SnowFlakeID) *User {
	u.iconImageID = &id
	return u
}

// Freeze ユーザーを凍結
func (u *User) Freeze() (*User, error) {
	if u.isFroze {
		return nil, errors.New("すでにユーザーは凍結されています")
	}

	u.isFroze = true
	return u, nil
}

// Unfreeze ユーザーを解凍(凍結解除)
func (u *User) Unfreeze() (*User, error) {
	if !u.isFroze {
		return nil, errors.New("ユーザーは凍結されていません")
	}

	u.isFroze = false
	return u, nil
}

// SetInboxURL ユーザーのInboxURLを設定
func (u *User) SetInboxURL(url string) (*User, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}

	u.inboxURL = url
	return u, nil
}

// SetOutboxURL ユーザーのOutboxURLを設定
func (u *User) SetOutboxURL(url string) (*User, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}

	u.outboxURL = url
	return u, nil
}

// SetFollowURL ユーザーのフォローURLを設定
func (u *User) SetFollowURL(url string) (*User, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}

	u.followURL = url
	return u, nil
}

// SetFollowerURL ユーザーのフォロワーURLを設定
func (u *User) SetFollowerURL(url string) (*User, error) {
	if len(url) <= 0 {
		return nil, errors.New("URLが短すぎます")
	}

	u.followersURL = url
	return u, nil
}

// SetSecretKey ユーザーの秘密鍵を設定
func (u *User) SetSecretKey(key string) (*User, error) {
	// ローカルユーザーにしか設定できない
	if !u.isLocalUser {
		return nil, errors.New("リモートユーザーに秘密鍵は設定できません")
	}

	if len(key) <= 0 {
		return nil, errors.New("鍵が短すぎます")
	}

	u.secretKey = &key
	return u, nil
}

// SetPublicKey ユーザーの公開鍵を設定
func (u *User) SetPublicKey(key string) (*User, error) {
	if len(key) <= 0 {
		return nil, errors.New("鍵が短すぎます")
	}

	u.publicKey = key
	return u, nil
}

// SetPassword ユーザーのパスワードを設定
func (u *User) SetPassword(pass string) (*User, error) {
	// ローカルユーザーにしか設定できない
	if !u.isLocalUser {
		return nil, errors.New("リモートユーザーにパスワードは設定できません")
	}

	if len(pass) <= 0 {
		return nil, errors.New("パスワードの文字数が短すぎます")
	}

	u.password = &pass
	return u, nil
}

func (u *User) SetUpdatedAt(now time.Time) (*User, error) {
	u.updatedAt = &now
	return u, nil
}

func (u *User) GetID() id.SnowFlakeID {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetDisplayName() string {
	return u.displayName
}

func (u *User) IsAdmin() bool {
	return u.role == AdminUserRole
}

func (u *User) GetInstanceID() id.SnowFlakeID {
	return u.instanceID
}

func (u *User) GetBio() *string {
	return u.bio
}

func (u *User) GetHeaderImageID() *id.SnowFlakeID {
	return u.headerImageID
}

func (u *User) GetIconImageID() *id.SnowFlakeID {
	return u.iconImageID
}

func (u *User) IsFroze() bool {
	return u.isFroze
}

func (u *User) GetInboxURL() string {
	return u.inboxURL
}

func (u *User) GetOutboxURL() string {
	return u.outboxURL
}

func (u *User) GetFollowURL() string {
	return u.followURL
}

func (u *User) GetFollowersURL() string {
	return u.followersURL
}

func (u *User) GetSecretKey() *string {
	return u.secretKey
}

func (u *User) GetPublicKey() string {
	return u.publicKey
}

func (u *User) GetPassword() *string {
	return u.password
}

func (u *User) IsLocalUser() bool {
	return u.isLocalUser
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetUpdatedAt() *time.Time {
	return u.updatedAt
}
