package controllers

import (
	"altastore/lib/database"
	"altastore/lib/utils"
	"altastore/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetProducts(c echo.Context) error {

	if utils.StringIsNotNumber(c.QueryParam("category")) {
		id, _ := strconv.Atoi(c.QueryParam("category"))
		product, err := database.GetProductsByCategoryId(id)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"status":  "fail",
				"message": err.Error(),
			})
		}

		res := models.Products_response{
			Code:    200,
			Status:  "Success",
			Message: "Success",
			Data:    product,
		}
		return c.JSON(http.StatusOK, res)
	} else if len(c.QueryParam("category")) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    400,
			Status:  "fail",
			Message: "invalid id supplied",
		})
	} else {
		product, err := database.GetProducts()

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"status":  "fail",
				"message": err.Error(),
			})
		}

		res := models.Products_response{
			Code:    200,
			Message: "Success",
			Status:  "Success",
			Data:    product,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func CreateProducts(c echo.Context) error {

	var post_body models.Products_post
	// c.Bind(&post_body)
	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Fail insert data",
			"status":  e.Error(),
		})
	}
	if e := models.Validate.Struct(post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Fail insert data",
			"status":  e.Error(),
		})
	}

	var product models.Products
	product.Categories_id = post_body.Categories_id
	product.Name = post_body.Name
	product.Description = post_body.Description
	product.Quantity = post_body.Quantity
	product.Price = post_body.Price
	product.Unit = post_body.Unit

	_, err := database.InsertProducts(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Fail insert data",
			"status":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    400,
		"message": "success Create products ",
		"data":    product,
	})
}

func UpdateProducts(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	}
	var post_body models.Products_update

	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": e.Error(),
		})
	}

	product, e := database.GetProductsById(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "Failed",
			"message": e.Error(),
		})
	}

	product.Name = utils.CompareStrings(product.Name, post_body.Name)
	product.Description = utils.CompareStrings(product.Description, post_body.Description)
	product.Unit = utils.CompareStrings(product.Unit, post_body.Unit)
	product.Quantity = utils.CompareId(product.Quantity, post_body.Quantity)
	product.Categories_id = utils.CompareId(product.Categories_id, post_body.Categories_id)
	product.Price = utils.CompareId(product.Price, post_body.Price)

	_, err := database.InsertProducts(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Fail insert data",
			"status":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success update products ",
		"data":    product,
	})

}

func DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	}

	err := database.DeleteProductsById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "Fail",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "product succesfully deleted",
	})
}
