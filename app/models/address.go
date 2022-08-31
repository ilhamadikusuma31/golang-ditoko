package models

import (
	"time"
)

type Address struct {
	ID         string `gorm:"size:36; not null; uniqueIndex; primary_key"`
	User       User
	UserID     string `gorm:"size:100; index"`
	Name       string `gorm:"size:100; not null"`
	IsPrimary  bool
	CityID     string `gorm:"size:100"`
	ProvinceID string `gorm:"size:100"`
	Address1   string `gorm:"size:255; not null"`
	Address2   string `gorm:"size:255; not null"`
	Phone      string `gorm:"size:100; not null"`
	Email      string `gorm:"size:100; not null"`
	PostCode   string `gorm:"size:100; not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
