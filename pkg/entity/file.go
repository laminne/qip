package entity

import "time"

type File struct {
	ID           string     `gorm:"primaryKey;not null"`
	FileName     string     `gorm:"not null;column:filename"`
	FilePath     *string    `gorm:"column:filepath"`
	FileURL      string     `gorm:"not null;column:fileurl"`
	ThumbnailURL *string    `gorm:"column:thumbnailurl"`
	Blurhash     string     `gorm:"not null"`
	IsNSFW       bool       `gorm:"not null;column:isnsfw"`
	MimeType     string     `gorm:"not null;column:mimetype"`
	CreatedAt    time.Time  `gorm:"precision:6;not null;column:createdat"`
	UpdatedAt    *time.Time `gorm:"precision:6;autoUpdateTime;column:updatedat"`

	UploaderID string `gorm:"column:uploaderid"`
	// ↓ 循環参照になるのでポインタにした
	Uploader *User   `gorm:"foreignKey:UploaderID"`
	PostID   *string `gorm:"column:postid"`
	Post     *Post   `gorm:"foreignKey:PostID"`
}
