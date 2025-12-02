package user

import (
	"orderApiStart/internal/product"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone  string          `json:"phone" gorm:"index"`
	Orders []product.Order `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewUser(phone string) *User {
	user := &User{
		Phone: phone,
	}
	return user
}
