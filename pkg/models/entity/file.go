package entity

import "time"

type File struct {
	ID        string    `bun:"id"`
	Host      string    `bun:"host"`
	Md5Hash   string    `bun:"md5Hash"`
	MimeType  string    `bun:"mimeType"`
	FileSize  int       `bun:"fileSize"`
	Url       string    `bun:"url"`
	IsNSFW    bool      `bun:"isNSFW"`
	BlurHash  string    `bun:"blurHash"`
	CreatedAt time.Time `bun:"createdAt"`

	UserID string `bun:"userID"`
}
