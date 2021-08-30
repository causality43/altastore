package database

import (
	"altastore/config"
	"altastore/models"
	"errors"
)

func InsertOrders(order models.Orders) (int, error) {
	if err := config.DB.Create(&order).Error; err != nil {
		return order.ID, err
	}
	return order.ID, nil
}

func GetOrderByCustomerId(customerId int) ([]models.Orders, error) {
	var order []models.Orders
	if rows := config.DB.Where("Customers_id = ?", customerId).Find(&order).RowsAffected; rows < 1 {
		return order, errors.New("user cart is empty")
	}
	return order, nil
}
