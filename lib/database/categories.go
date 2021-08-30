package database

import (
	"altastore/config"
	"altastore/models"
	"errors"
)

func GetCategories() ([]models.Categories, error) {
	var categories []models.Categories
	if err := config.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func InsertCategories(categories models.Categories) (interface{}, error) {

	if err := config.DB.Save(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoriesById(id int) (models.Categories, error) {
	var categories models.Categories

	if rows := config.DB.Find(&categories, id).RowsAffected; rows < 1 {
		err := errors.New("categories not found")
		return categories, err
	}

	return categories, nil
}

func DeleteCategoriesById(id int) error {

	rows := config.DB.Delete(&models.Categories{}, id).RowsAffected
	if rows == 0 {
		err := errors.New("categories to be deleted is not found")
		return err
	}
	return nil

}
