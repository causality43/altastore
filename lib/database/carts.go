package database

import (
	"altastore/config"
	"altastore/models"
	"errors"
)

func InsertCarts(cart models.Carts) (interface{}, error) {

	if err := config.DB.Save(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func GetCartsById(id int) (models.Carts, error) {
	var carts models.Carts

	if rows := config.DB.Find(&carts, id).RowsAffected; rows < 1 {
		err := errors.New("carts not found")
		return carts, err
	}
	return carts, nil
}

func GetCartsIdFromUser(customerId int) (int, error) {
	var carts models.Carts
	if rows := config.DB.Where("Customers_id = ?", customerId).Find(&carts).RowsAffected; rows < 1 {
		return -1, errors.New("carts not found")
	}
	return carts.ID, nil
}
