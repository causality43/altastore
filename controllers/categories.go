package controllers

import (
	"altastore/lib/database"
	"altastore/lib/utils"
	"altastore/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetCategories(c echo.Context) error {
	categories, err := database.GetCategories()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": err.Error(),
		})
	}

	res := models.Categories_response{
		Code:    200,
		Message: "Success",
		Status:  "Success",
		Data:    categories,
	}
	return c.JSON(http.StatusOK, res)
}

func CreateCategories(c echo.Context) error {
	var post_body models.Categories_post

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

	var categories models.Categories
	categories.Name = post_body.Name
	categories.Description = post_body.Description

	_, err := database.InsertCategories(categories)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Categories_response_single{
		Code:    200,
		Status:  "success",
		Message: "success insert categories",
		Data:    categories,
	})
}

func UpdateCategories(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	}
	var post_body models.Categories_update

	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "fail",
			"message": e.Error(),
		})
	}

	categories, e := database.GetCategoriesById(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "fail",
			"message": e.Error(),
		})
	}

	categories.Name = utils.CompareStrings(categories.Name, post_body.Name)
	categories.Description = utils.CompareStrings(categories.Description, post_body.Description)

	_, err := database.InsertCategories(categories)
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
		"message": "success update categories ",
		"data":    categories,
	})
}

func DeleteCategories(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	}

	err := database.DeleteCategoriesById(id)
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
		"message": "category succesfully deleted",
	})
}
