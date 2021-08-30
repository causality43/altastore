package models

import (
	"time"
)

type Customers struct {
	ID                  int    `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Username            string `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email               string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password            string `gorm:"type:varchar(255);not null" json:"password"`
	Address             string `gorm:"type:varchar(255);" json:"address"`
	Bank_name           string `gorm:"type:varchar(255);" json:"bank_name"`
	Bank_account_number int    `gorm:"type:bigint(20);default:0;" json:"bank_account_number"`
	CreatedAt           time.Time
}

type Customers_register struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Customers_login struct {
	ID    int    `gorm:"primarykey;" json:"id"`
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token"`
}

type Customers_update struct {
	Username            string `json:"username"`
	Email               string `json:"email" validate:"required,email"`
	Address             string `json:"address"`
	Bank_name           string `json:"bank_name"`
	Bank_account_number int    `json:"bank_account_number"`
}

type Customers_response struct {
	Code    string
	Message string
	Status  string
}
