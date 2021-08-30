package models

import "time"

type Orders struct {
	ID           int       `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Customers_id int       `gorm:"not null" json:"customers_id"`
	Couriers_id  int       `gorm:"not null" json:"couriers_id"`
	Address      string    `gorm:"type:varchar(255);not null" json:"address"`
	Customer     Customers `gorm:"foreignKey:ID"`
	Courier      Couriers  `gorm:"foreignkey:ID"`
	CreatedAt    time.Time
}

type Orders_post struct {
	Customers_id       int     `json:"customers_id" validate:"required"`
	Couriers_id        int     `json:"couriers_id" validate:"required"`
	Payment_method     string  `json:"payment_method" validate:"required"`
	Payment_start_date string  `json:"payment_start_date" validate:"required"`
	Payment_end_date   string  `json:"payment_end_date" validate:"required"`
	Payment_status     string  `json:"payment_status" validate:"required"`
	Payment_amount     float32 `json:"payment_amount" validate:"required"`
}

type Orders_arr_response struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Status  string   `json:"status"`
	Data    []Orders `json:"data"`
}
