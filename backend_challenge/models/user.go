package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model `json:"-"`
	ID uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Email string  `json:"email" gorm:"not null;unique"`
	Password string `json:"-"`
	DisplayConfig  string `gorm:"type:json" json:"display_config"` 
}


type UserInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
	DisplayConfig  string `gorm:"type:json" json:"display_config"`     
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHashed)
	return nil
}
