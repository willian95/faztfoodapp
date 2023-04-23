package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string  `gorm:"type:text; FULLTEXT" json:"name"`
	Picture     string  `gorm:"type:varchar(250)" json:"picture"`
	Description string  `gorm:"type:text; FULLTEXT" json:"description"`
	Subtitle    string  `gorm:"type:text; FULLTEXT" json:"subtitle"`
	Price       float64 `gorm:"type:double" json:"price"`
	Rating      float64 `gorm:"type:double" json:"rating"`
}
