package controllers

import (
	"altastore/lib/database"
	"altastore/lib/utils"
	"altastore/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateOrders(c echo.Context) error {
	var post_body models.Orders_post

	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": e.Error(),
		})
	}
	if e := models.Validate.Struct(post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Error",
			"message": e.Error(),
		})
	}

	var newOrder models.Orders
	newOrder.Customers_id = post_body.Customers_id
	newOrder.Couriers_id = post_body.Couriers_id
	newOrder.Address, _ = database.GetCustomersAddress(post_body.Customers_id)

	//save order
	orderID, err := database.InsertOrders(newOrder)
	//create payment
	database.InsertPaymentsWithOrderId(post_body, orderID)
	//pindahin semua item dari keranjang user ke checkout items
	cartItemsArr, _ := database.ExtractCartItemsFromUser(post_body.Customers_id)
	database.InsertCheckoutItems(cartItemsArr, orderID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.SuccessResponse{
		Code:    200,
		Status:  "success",
		Message: "success insert order",
	})
}

func GetOrder(c echo.Context) error {
	userId := c.QueryParam("user")
	if !utils.StringIsNotNumber(userId) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	}
	id, _ := strconv.Atoi(userId)
	orderArr, _ := database.GetOrderByCustomerId(id)
	return c.JSON(http.StatusOK, models.Orders_arr_response{
		Code:    200,
		Status:  "success",
		Message: "success get order",
		Data:    orderArr,
	})

}
