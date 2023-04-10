package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	Slug string `gorm:"unique,not null" uri:"slug"`
	Full string `gorm:"not null"        uri:"full"`
}
