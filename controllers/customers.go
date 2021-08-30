package controllers

import (
	"altastore/config"
	"altastore/lib/database"
	"altastore/lib/utils"
	"altastore/middlewares"
	"altastore/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterCustomersController(c echo.Context) error {
	var customerModel models.Customers_register

	if e := c.Bind(&customerModel); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Error registration customers",
			"status":  e.Error(),
		})
	}

	// cek validation struct
	if e := models.Validate.Struct(customerModel); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Error registration customers",
			"status":  e.Error(),
		})
	}

	// Generate "hash" to store from user password
	// hash, _ := HashPassword(customerModel.Password)

	var customer models.Customers
	customer.Username = customerModel.Username
	customer.Email = customerModel.Email
	customer.Password = customerModel.Password

	_, err := database.InsertCustomers(customer)

	// cek before insert to database
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Error registration customers",
			"status":  err.Error(),
		})
	}

	//create cart untuk customer
	newCustomer, _ := database.GetCustomersByName(customer.Username)
	newCart := models.Carts{
		Customers_id: newCustomer.ID,
		Status:       "empty",
	}
	_, e := database.InsertCarts(newCart)

	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Error registration user",
			"status":  e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success register customers",
		"status":  "success",
	})
}

func LoginCustomersController(c echo.Context) error {
	var customerData models.Customers
	var customerLogin models.Customers_login
	var err error
	c.Bind(&customerData)

	if err = config.DB.Where("email = ? AND password = ?", customerData.Email, customerData.Password).First(&customerData).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Invalid email or password",
			"status":  err.Error(),
		})
	}

	customerLogin.Token, err = middlewares.CreateToken(int(customerData.ID))
	customerLogin.Email = customerData.Email
	customerLogin.ID = customerData.ID

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Error invalid JWT",
			"status":  "Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success login customer",
		"status":  "success",
		"data":    customerLogin,
	})
}

func UpdateProfileCustomersController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id customers",
		})
	}

	var post_body models.Customers_update

	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "fail",
			"message": e.Error(),
		})
	}

	customers, e := database.GetCustomersByid(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "fail",
			"message": e.Error(),
		})
	}

	customers.Username = utils.CompareStrings(customers.Username, post_body.Username)
	customers.Email = utils.CompareStrings(customers.Email, post_body.Email)
	customers.Address = utils.CompareStrings(customers.Address, post_body.Address)
	customers.Bank_name = utils.CompareStrings(customers.Bank_name, post_body.Bank_name)
	customers.Bank_account_number = post_body.Bank_account_number

	_, err := database.InsertCustomers(customers)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success update profil customer",
		"data":    customers,
	})

}
