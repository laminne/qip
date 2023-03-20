package entity

import "time"

type User struct {
	ID           string `gorm:"primaryKey;not null"`
	Name         string `gorm:"not null"`
	DisplayName  string `gorm:"column:displayname"`
	Role         int    `gorm:"not null"`
	Bio          *string
	IsFroze      bool    `gorm:"column:isfloze"`
	InboxURL     string  `gorm:"column:inboxurl"`
	OutboxURL    string  `gorm:"column:outboxurl"`
	FollowURL    string  `gorm:"column:followurl"`
	FollowersURL string  `gorm:"column:followersurl"`
	SecretKey    *string `gorm:"column:secretkey"`
	PublicKey    string  `gorm:"column:publickey"`
	Password     *string
	IsLocalUser  bool       `gorm:"column:islocaluser"`
	CreatedAt    time.Time  `gorm:"column:createdat"`
	UpdatedAt    *time.Time `gorm:"column:updatedat"`

	InstanceID string   `gorm:"column:instanceid"`
	Instance   Instance `gorm:"foreignKey:InstanceID"`

	HeaderImageID *string `gorm:"column:headerimageid"`
	HeaderImage   File    `gorm:"foreignKey:HeaderImageID"`

	IconImageID *string `gorm:"column:iconimageid"`
	IconImage   File    `gotm:"foreignKey:IconImageID"`
}
