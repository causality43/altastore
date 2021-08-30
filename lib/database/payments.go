package database

import (
	"altastore/config"
	"altastore/models"
	"errors"
	"time"
)

func InsertPaymentsWithOrderId(order models.Orders_post, order_id int) error {
	var newPayment models.Payments
	timeFormat := "2006-01-02 15:04:05"
	newPayment.Order_id = order_id
	newPayment.Payment_amount = order.Payment_amount
	newPayment.Payment_method = order.Payment_method
	newPayment.Payment_status = order.Payment_status
	newPayment.Payment_start_date, _ = time.Parse(timeFormat, order.Payment_start_date)
	newPayment.Payment_end_date, _ = time.Parse(timeFormat, order.Payment_end_date)

	if err := config.DB.Save(&newPayment).Error; err != nil {
		return err
	}
	return nil
}

func InsertPayments(payment models.Payments) error {

	if err := config.DB.Save(&payment).Error; err != nil {
		return err
	}
	return nil
}

func GetAllPayments() []models.Payments {
	var payment []models.Payments
	config.DB.Find(&payment)
	return payment
}

func GetPaymentByOrderId(orderId int) (models.Payments, error) {
	var payment models.Payments
	if err := config.DB.Where("Order_id = ?", orderId).Find(&payment).Error; err != nil {
		return payment, err
	}
	return payment, nil
}

func GetPaymentById(id int) (models.Payments, error) {
	var payment models.Payments
	if rows := config.DB.Where("ID = ?", id).Find(&payment).RowsAffected; rows < 1 {
		return payment, errors.New("no payment found")
	}
	return payment, nil
}
