package controllers

import (
	"altastore/lib/database"
	"altastore/lib/utils"
	"altastore/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateCartitems(c echo.Context) error {
	var post_body models.Cartitems_Post

	// c.Bind(&post_body)
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

	var cartitems models.Cartitems
	cartitems.Carts_id = *post_body.Carts_id
	cartitems.Products_id = *post_body.Products_id
	cartitems.Quantity = *post_body.Quantity

	_, err := database.InsertCartitems(cartitems)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Cartitems_response_single{
		Code:    200,
		Status:  "success",
		Message: "success insert cartitems",
		Data:    cartitems,
	})
}

func GetCartitemsByCartId(c echo.Context) error {
	if utils.StringIsNotNumber(c.QueryParam("cart")) {
		id, _ := strconv.Atoi(c.QueryParam("cart"))
		cartItems, err := database.GetCartitemsByCartsId(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"status":  "fail",
				"message": err.Error(),
			})
		}
		if len(cartItems) == 0 {
			res := models.Cartitems_response{
				Code:    200,
				Status:  "Success",
				Message: "Success",
				Data:    cartItems,
			}
			return c.JSON(http.StatusOK, res)
		}
		cartResponse := database.ConvertIntoCartResponse(cartItems)

		res := models.CartItems_response_detail{
			Code:    200,
			Status:  "success",
			Message: "Success Get Cartitems",
			Data:    cartResponse,
		}
		return c.JSON(http.StatusOK, res)
	} else if len(c.QueryParam("cart")) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    400,
			Status:  "fail",
			Message: "invalid id supplied",
		})
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    400,
			Status:  "fail",
			Message: "invalid method",
		})
	}
}

func UpdateCartitems(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	}
	var post_body models.Cartitems_Update

	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "fail",
			"message": e.Error(),
		})
	}

	cartitems, e := database.GetCartitemsById(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "fail",
			"message": e.Error(),
		})
	}

	cartitems.Quantity = *post_body.Quantity

	_, err := database.InsertCartitems(cartitems)
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
		"message": "success update cartitems ",
		"data":    cartitems,
	})
}

func DeleteCartitems(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	}

	err := database.DeleteCartitemsById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "Fail",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Code:    200,
		Status:  "success",
		Message: "cartitems succesfully deleted",
	})
}
