package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone string `json:"phone" gorm:"index"`
}

func NewUser(phone string) *User {
	user := &User{
		Phone: phone,
	}
	return user
}
