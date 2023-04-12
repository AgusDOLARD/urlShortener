package models

import "gorm.io/gorm"

// URL url model
type URL struct {
	gorm.Model
	Slug string `gorm:"unique,not null" json:"slug"`
	Full string `gorm:"not null"        json:"full" validate:"required,url"`
}
