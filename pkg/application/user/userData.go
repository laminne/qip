package user

import (
	"time"

	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/utils/id"
)

type UserData struct {
	id            id.SnowFlakeID
	name          string
	displayName   string
	role          domain.UserRole
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

func NewUserData(u domain.User) *UserData {
	role := 0
	if u.IsAdmin() {
		role = 1
	}
	return &UserData{
		id:            u.GetID(),
		name:          u.GetName(),
		displayName:   u.GetDisplayName(),
		role:          role,
		instanceID:    u.GetInstanceID(),
		bio:           u.GetBio(),
		headerImageID: u.GetHeaderImageID(),
		iconImageID:   u.GetIconImageID(),
		isFroze:       u.IsFroze(),
		inboxURL:      u.GetInboxURL(),
		outboxURL:     u.GetOutboxURL(),
		followURL:     u.GetFollowURL(),
		followersURL:  u.GetFollowersURL(),
		secretKey:     u.GetSecretKey(),
		publicKey:     u.GetPublicKey(),
		password:      u.GetPassword(),
		isLocalUser:   u.IsLocalUser(),
		createdAt:     u.GetCreatedAt(),
		updatedAt:     u.GetUpdatedAt(),
	}
}

func (u UserData) Id() id.SnowFlakeID {
	return u.id
}

func (u UserData) Name() string {
	return u.name
}

func (u UserData) DisplayName() string {
	return u.displayName
}

func (u UserData) Role() domain.UserRole {
	return u.role
}

func (u UserData) InstanceID() id.SnowFlakeID {
	return u.instanceID
}

func (u UserData) Bio() *string {
	return u.bio
}

func (u UserData) HeaderImageID() *id.SnowFlakeID {
	return u.headerImageID
}

func (u UserData) IconImageID() *id.SnowFlakeID {
	return u.iconImageID
}

func (u UserData) IsFroze() bool {
	return u.isFroze
}

func (u UserData) InboxURL() string {
	return u.inboxURL
}

func (u UserData) OutboxURL() string {
	return u.outboxURL
}

func (u UserData) FollowURL() string {
	return u.followURL
}

func (u UserData) FollowersURL() string {
	return u.followersURL
}

func (u UserData) SecretKey() *string {
	return u.secretKey
}

func (u UserData) PublicKey() string {
	return u.publicKey
}

func (u UserData) Password() *string {
	return u.password
}

func (u UserData) IsLocalUser() bool {
	return u.isLocalUser
}

func (u UserData) CreatedAt() time.Time {
	return u.createdAt
}

func (u UserData) UpdatedAt() *time.Time {
	return u.updatedAt
}
