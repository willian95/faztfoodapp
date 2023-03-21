package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `gorm: "not null"`
	Password string `gorm: "not null"`
	Email    string `gorm: "not null"; unique_index`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	if u.Password != "" {
		hash, err := MakePassword(u.Password)
		if err != nil {
			return err
		}
		tx.Statement.SetColumn("Password", hash)
	}

	return
}

// MakePassword : Encrypt user password
func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
