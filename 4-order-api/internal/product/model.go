package product

import (
	"github.com/lib/pq"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images"`
	Order       []Order        `gorm:"many2many:order_product;"`
}

type Order struct {
	gorm.Model
	UserID   uint      `json:"userId"`
	Products []Product `gorm:"many2many:order_product;"`
}

func NewOrder(userId uint, products []Product) *Order {
	order := &Order{
		UserID:   userId,
		Products: products,
	}
	return order
}

func NewProduct(name, description string) *Product {
	product := &Product{
		Name:        name,
		Description: description,
	}

	return product
}
