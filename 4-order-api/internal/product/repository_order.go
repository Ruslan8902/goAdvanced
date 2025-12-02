package product

import "orderApiStart/pkg/db"

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{
		Database: database,
	}
}

func (repo *OrderRepository) Create(order *Order) (*Order, error) {
	result := repo.Database.DB.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (repo *OrderRepository) GetByOrderId(userId uint, orderId uint) (*[]Order, error) {
	var orders []Order
	result := repo.Database.DB.Find(&orders, "UserID = ? AND OrderID = ?", userId, orderId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orders, nil
}

func (repo *OrderRepository) GetByUserId(id uint) (*[]Order, error) {
	var orders []Order
	result := repo.Database.DB.Find(&orders, "UserID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orders, nil
}
