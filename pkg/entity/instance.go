package entity

import "time"

type Instance struct {
	ID              string     `gorm:"primaryKey;not null"`
	Name            string     `gorm:"not null"`
	SoftwareName    string     `gorm:"not null;column:softwarename"`
	SoftwareVersion string     `gorm:"not null;column:softwareversion"`
	Host            string     `gorm:"unique"`
	Description     string     `gorm:"not null"`
	State           int        `gorm:"not null"`
	CreatedAt       time.Time  `gorm:"precision:6;not null;column:createdat"`
	UpdatedAt       *time.Time `gorm:"precision:6;autoUpdateTime;column:updatedat"`
}
