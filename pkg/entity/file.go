package entity

import "time"

type File struct {
	ID           string `gormRepository:"primaryKey;not null"`
	FileName     string `gormRepository:"not null"`
	FilePath     *string
	FileURL      string `gormRepository:"not null"`
	ThumbnailURL *string
	Blurhash     string     `gormRepository:"not null"`
	IsNSFW       bool       `gormRepository:"not null"`
	MimeType     string     `gormRepository:"not null"`
	CreatedAt    time.Time  `gormRepository:"precision:6;not null"`
	UpdatedAt    *time.Time `gormRepository:"precision:6;autoUpdateTime"`
}
