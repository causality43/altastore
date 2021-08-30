package models

type Couriers struct {
	ID   int    `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"type:varchar(255);unique;not null" json:"name"`
}
