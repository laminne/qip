package entity

import "time"

type Post struct {
	ID         string    `gorm:"primaryKey;not null"`
	Body       string    `gorm:"type:text;not null"`
	Visibility int       `gorm:"not null"`
	AuthorID   string    `gorm:"not null;column:authorid"`
	CreatedAt  time.Time `gorm:"column:createdat"`
}
