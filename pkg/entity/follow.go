package entity

import "time"

type Follow struct {
	UserID    string    `gorm:"column:userid;not null"`
	TargetID  string    `gorm:"column:targetid;not null"`
	CreatedAt time.Time `gorm:"precision:6;not null;column:createdat"`
}
