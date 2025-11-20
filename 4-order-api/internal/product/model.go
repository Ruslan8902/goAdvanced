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
}

func NewProduct(name, description string) *Product {
	product := &Product{
		Name:        name,
		Description: description,
	}

	return product
}
