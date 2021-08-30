package models

import (
	// "gorm.io/gorm"
	"time"
)

type Categories struct {
	ID          int    `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Categories_response struct {
	Code    int          `json:"code"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    []Categories `json:"data"`
}

type Categories_response_single struct {
	Code    int        `json:"code"`
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    Categories `json:"data"`
}

type Categories_post struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type Categories_update struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
