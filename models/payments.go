package models

import (
	"time"
)

type Payments struct {
	ID                 int       `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Order_id           int       `gorm:"not null" json:"order_id"`
	Payment_method     string    `gorm:"type:varchar(255);not null" json:"payment_method"`
	Payment_start_date time.Time `gorm:"not null" json:"payment_start_date"`
	Payment_end_date   time.Time `gorm:"not null" json:"payment_end_date"`
	Payment_status     string    `gorm:"type:varchar(255);not null" json:"payment_status"`
	Payment_amount     float32   `gorm:"type:float;not null" json:"payment_amount"`
	Order              Orders    `gorm:"foreignkey:ID" json:"-"`
}

type Payments_update struct {
	Payment_status string `json:"payment_status" validate:"required"`
}
