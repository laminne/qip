package entity

import "time"

type File struct {
	ID           string `gorm:"primaryKey;not null"`
	FileName     string `gorm:"not null"`
	FilePath     *string
	FileURL      string `gorm:"not null"`
	ThumbnailURL *string
	Blurhash     string     `gorm:"not null"`
	IsNSFW       bool       `gorm:"not null"`
	MimeType     string     `gorm:"not null"`
	CreatedAt    time.Time  `gorm:"precision:6;not null"`
	UpdatedAt    *time.Time `gorm:"precision:6;autoUpdateTime"`

	UploaderID string `gorm:"column:uploaderid"`
	// ↓ 循環参照になるのでポインタにした
	Uploader *User   `gorm:"foreignKey:UploaderID"`
	PostID   *string `gorm:"column:postid"`
	Post     *Post   `gorm:"foreignKey:PostID"`
}
