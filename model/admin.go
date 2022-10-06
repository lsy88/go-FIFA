package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name              string `json:"name" gorm:"type:varchar;column:name"`
	EncryptedPassword string `json:"encrypted_password" gorm:"type:varchar;column:encrypted_password'"`
	Phone             string `json:"phone" gorm:"type:varchar;column:phone"`
	State             string `json:"state" gorm:"type:varchar;column:state"`
}

func (a *Admin) PasswordCheck(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
