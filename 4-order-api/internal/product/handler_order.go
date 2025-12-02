package product

import (
	"net/http"
	"orderApiStart/configs"
	"orderApiStart/middleware"
	"orderApiStart/pkg/req"
	"orderApiStart/pkg/res"
	"strconv"
)

type OrderHandlerDeps struct {
	OrderRepository *OrderRepository
	Config          *configs.Config
}

type OrderHandler struct {
	OrderRepository *OrderRepository
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepository: deps.OrderRepository,
	}

	router.Handle("POST /order", middleware.IsAuthed(handler.Create(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuthed(handler.GetMyOrders(), deps.Config))
	router.Handle("GET /order/{id}", middleware.IsAuthed(handler.Get(), deps.Config))
}

func (handler *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[OrderCreateRequest](&w, r)
		if err != nil {
			return
		}
		order := &Order{UserID: body.UserID,
			Products: body.Products,
		}

		createdOrder, err := handler.OrderRepository.Create(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdOrder, 201)
	}
}

func (handler *OrderHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userId, ok := r.Context().Value(middleware.ContextUserIDKey).(uint)
		if !ok {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		orders, err := handler.OrderRepository.GetByOrderId(uint(userId), uint(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		res.Json(w, orders, 200)
	}
}

func (handler *OrderHandler) GetMyOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value(middleware.ContextUserIDKey).(uint)
		if !ok {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		orders, err := handler.OrderRepository.GetByUserId(uint(userId))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		res.Json(w, orders, 200)
	}
}
