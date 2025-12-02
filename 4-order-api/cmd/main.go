package main

import (
	"fmt"
	"net/http"
	"orderApiStart/configs"
	"orderApiStart/internal/auth"
	"orderApiStart/internal/product"
	"orderApiStart/internal/user"
	"orderApiStart/middleware"
	"orderApiStart/migrations"
	"orderApiStart/pkg/db"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	migrations.AutoMigrate()

	router := http.NewServeMux()

	productRepository := product.NewProductRepository(db)
	sesseionRepository := auth.NewSessionRepository(db)
	userRepository := user.NewUserRepository(db)
	orderRepository := product.NewOrderRepository(db)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
		Config:            conf,
	})

	product.NewOrderHandler(router, product.OrderHandlerDeps{
		OrderRepository: orderRepository,
		Config:          conf,
	})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		SessionRepository: sesseionRepository,
		UserRepository:    userRepository,
		Config:            conf,
	})

	return middleware.Logging(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
