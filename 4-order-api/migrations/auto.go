package migrations

import (
	"orderApiStart/internal/auth"
	"orderApiStart/internal/product"
	"os"
	"os/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AutoMigrate() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&product.Product{}, &auth.Session{}, &user.User{})
}
