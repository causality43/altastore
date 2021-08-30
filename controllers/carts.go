package controllers

import (
	"altastore/lib/database"
	"altastore/lib/utils"
	"altastore/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func UpdateCarts(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	}
	var post_body models.Carts_update

	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "fail",
			"message": e.Error(),
		})
	}

	if e := models.Validate.Struct(post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Fail insert data",
			"status":  e.Error(),
		})
	}

	carts, e := database.GetCartsById(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "fail",
			"message": e.Error(),
		})
	}

	carts.Status = post_body.Status

	_, err := database.InsertCarts(carts)
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
		"data":    carts,
	})
}
