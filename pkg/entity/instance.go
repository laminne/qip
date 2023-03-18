package entity

import "time"

type Instance struct {
	ID              string     `gormRepository:"primaryKey;not null"`
	Name            string     `gormRepository:"not null"`
	SoftwareName    string     `gormRepository:"not null"`
	SoftwareVersion string     `gormRepository:"not null"`
	Host            string     `gormRepository:"unique"`
	Description     string     `gormRepository:"not null"`
	State           int        `gormRepository:"not null"`
	CreatedAt       time.Time  `gormRepository:"precision:6;not null"`
	UpdatedAt       *time.Time `gormRepository:"precision:6;autoUpdateTime"`
}
