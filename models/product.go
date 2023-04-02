package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string  `gorm:"type:varchar(250)" json:"name"`
	Picture     string  `gorm:"type:varchar(250)" json:"picture"`
	Description string  `gorm:"type:text" json:"description"`
	Subtitle    string  `gorm:"type:varchar(250)" json:"subtitle"`
	Rating      float64 `gorm:"type:double" json:"rating"`
}
