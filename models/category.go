package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	Name    string `gorm:"type:varchar(250)" json:"name"`
	Picture string `gorm:"type:varchar(250)" json:"picture"`
	Slug    string `gorm:"type:varchar(250)" json:"slug"`
}
