package database

import (
	"altastore/config"
	"altastore/models"
)

func InsertCheckoutItems(cartItems []models.Cartitems, orderId int) error {

	for _, item := range cartItems {

		var checkout_items models.Checkout_items
		checkout_items.Order_id = orderId
		checkout_items.Products_id = item.Products_id
		checkout_items.Quantity = item.Quantity

		if err := config.DB.Save(&checkout_items).Error; err != nil {
			return err
		}
	}

	return nil
}

func GetCheckOutItemByOrderId(id int) ([]models.Checkout_items, error) {
	var items []models.Checkout_items
	if e := config.DB.Where("Order_id = ?", id).Find(&items).Error; e != nil {
		return items, e
	}
	return items, nil
}
