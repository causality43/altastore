package database

import (
	"altastore/config"
	"altastore/models"
)

func InsertCustomers(customer models.Customers) (interface{}, error) {
	if err := config.DB.Save(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomersByid(id int) (models.Customers, error) {
	var customer models.Customers
	var empty models.Customers

	if err := config.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		return empty, err
	}
	if err := config.DB.Find(&customer, id).Error; err != nil {
		return empty, err
	}

	return customer, nil
}

func GetCustomersByEmail(email string) (models.Customers, error) {
	var customer models.Customers
	var empty models.Customers

	if err := config.DB.Where("email = ?", email).First(&customer).Error; err != nil {
		return empty, err
	}
	if err := config.DB.Find(&customer, email).Error; err != nil {
		return empty, err
	}

	return customer, nil
}

func GetCustomersByName(name string) (models.Customers, error) {
	var customer models.Customers
	if e := config.DB.Where("Username = ?", name).Find(&customer).Error; e != nil {
		return customer, e
	}

	return customer, nil
}

func GetCustomersAddress(customerId int) (string, error) {
	var customer models.Customers
	if e := config.DB.Where("ID = ?", customerId).Find(&customer).Error; e != nil {
		return " ", e
	}
	return customer.Address, nil
}
