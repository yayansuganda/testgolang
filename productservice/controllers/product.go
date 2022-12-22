package controllers

import (
	"net/http"
	"productservice/models"

	"github.com/labstack/echo"
)

func StoreProduct(c echo.Context) error {
	product := new(models.Product)
	c.Bind(product)
	response := new(models.Response)
	if product.CreateProduct() != nil { // method create  product
		response.ErrorCode = 10
		response.Message = "Gagal create data Product"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses create data product"
		response.Data = *product
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateProduct(c echo.Context) error {
	product := new(models.Product)
	c.Bind(product)
	response := new(models.Response)
	if product.UpdateProduct(c.Param("id")) != nil { // method update  product
		response.ErrorCode = 10
		response.Message = "Gagal update data product"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses update data product"
		response.Data = *product
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteProduct(c echo.Context) error {
	product := new(models.Product)
	// product, _ := models.GetOneById(c.Param("id")) // method get by email
	response := new(models.Response)

	if product.DeleteProduct(c.Param("id")) != nil { // method update product
		response.ErrorCode = 10
		response.Message = "Gagal menghapus data product"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses menghapus data product"
	}
	return c.JSON(http.StatusOK, response)
}

// GetProduct
// @Summary Fetch a list of all product.
// @Description Fetch a list of all product.
// @Tags Product
// @Accept */*
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Security Bearer Authentication
// @Produce json
// @Success 200
// @Failure 400
// @Router /product [get]
func GetProduct(c echo.Context) error {
	response := new(models.Response)
	product, err := models.GetProductAll() // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal melihat data products"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses melihat data products"
		response.Data = product
	}
	return c.JSON(http.StatusOK, response)
}

func GetSingleProduct(c echo.Context) error {
	response := new(models.Response)
	product, err := models.GetProductOneById(c.Param("id")) // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal melihat data products"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses melihat data products"
		response.Data = product
	}
	return c.JSON(http.StatusOK, response)
}
