package config

import (
	"altastore/models"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error reading .env file")
	}

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DBNAME"],
	)

	DB, err = gorm.Open(mysql.Open(mysqlCredentials), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Customers{})
	DB.AutoMigrate(&models.Categories{})
	DB.AutoMigrate(&models.Products{})
	DB.AutoMigrate(&models.Carts{})
	DB.AutoMigrate(&models.Cartitems{})
	DB.AutoMigrate(&models.Checkout_items{})
	DB.AutoMigrate(&models.Orders{})
	DB.AutoMigrate(&models.Couriers{})
	DB.AutoMigrate(&models.Payments{})
}
