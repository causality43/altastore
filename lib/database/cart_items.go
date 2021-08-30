package database

import (
	"altastore/config"
	"altastore/models"
	"errors"
)

func GetCartitemsByCartsId(cart_id int) ([]models.Cartitems, error) {
	var cartitems []models.Cartitems

	if e := config.DB.Where("carts_id = ?", cart_id).Find(&cartitems).Error; e != nil {
		return cartitems, e
	}

	return cartitems, nil
}

func InsertCartitems(cartitems models.Cartitems) (interface{}, error) {

	if err := config.DB.Save(&cartitems).Error; err != nil {
		return nil, err
	}
	return cartitems, nil
}

func GetCartitemsById(id int) (models.Cartitems, error) {
	var cartitems models.Cartitems

	if rows := config.DB.Find(&cartitems, id).RowsAffected; rows < 1 {
		err := errors.New("cartitems not found")
		return cartitems, err
	}

	return cartitems, nil
}

func DeleteCartitemsById(id int) error {

	rows := config.DB.Delete(&models.Cartitems{}, id).RowsAffected
	if rows == 0 {
		err := errors.New("cartitems to be deleted is not found")
		return err
	}
	return nil
}

func ExtractCartItemsFromUser(userId int) ([]models.Cartitems, error) {
	userCartId, _ := GetCartsIdFromUser(userId)
	//ambil semua item cart milik user
	var cartItems []models.Cartitems
	if rows := config.DB.Where("Carts_id = ?", userCartId).Find(&cartItems).RowsAffected; rows < 1 {
		return cartItems, errors.New("user cart is empty")
	}

	//kosongkan semua item yg sudah diambil
	for _, item := range cartItems {
		config.DB.Delete(&models.Cartitems{}, item.ID)
	}

	return cartItems, nil
}

func ConvertIntoCartResponse(cartItems []models.Cartitems) []models.CartItems_response_user {
	cartArr := []models.CartItems_response_user{}
	for _, item := range cartItems {
		var product models.Products
		config.DB.Where("ID = ?", item.Products_id).Find(&product)

		a := models.CartItems_response_user{
			ID:          item.ID,
			Carts_id:    item.Carts_id,
			Name:        product.Name,
			Price:       product.Price,
			Quantity:    item.Quantity,
			Products_id: item.Products_id,
		}
		cartArr = append(cartArr, a)
	}
	return cartArr
}
