package database

import (
	"altastore/config"
	"altastore/models"
	"errors"
)

func GetProducts() ([]models.Products, error) {
	var product []models.Products
	if err := config.DB.Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func GetProductsById(id int) (models.Products, error) {
	var product models.Products
	var empty models.Products

	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return empty, err
	}
	if err := config.DB.Find(&product, id).Error; err != nil {
		return empty, err
	}

	return product, nil
}

func GetProductsByCategoryId(id int) ([]models.Products, error) {
	var product []models.Products
	if e := config.DB.Where("categories_id = ?", id).Find(&product).Error; e != nil {
		return product, e
	}
	return product, nil
}

func GetProductUser([]models.Products, error) {

}

func InsertProducts(product models.Products) (interface{}, error) {

	if err := config.DB.Save(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func DeleteProductsById(id int) error {

	rows := config.DB.Delete(&models.Products{}, id).RowsAffected
	if rows == 0 {
		err := errors.New("product to be deleted is not found")
		return err
	}
	return nil

}
