package controllers

import (
	"net/http"
	"orderservice/models"

	"github.com/labstack/echo"
)

func StoreOrder(c echo.Context) error {
	order := new(models.Order)
	product := new(models.Product)
	c.Bind(order)
	response := new(models.Response)
	if order.CreateOrder() != nil { // method create  order
		response.ErrorCode = 10
		response.Message = "Gagal create data user"
	} else {
		productid := *&order.ProductId
		jumlah := *&order.Jumlah

		product.UpdateStok(productid, jumlah)
		response.ErrorCode = 0
		response.Message = "Sukses create data order"
		response.Data = *order
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateOrder(c echo.Context) error {
	order := new(models.Order)
	c.Bind(order)
	response := new(models.Response)
	if order.UpdateOrder(c.Param("id")) != nil { // method update  order
		response.ErrorCode = 10
		response.Message = "Gagal update data order"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses update data order"
		response.Data = *order
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteOrder(c echo.Context) error {
	order := new(models.Order)
	// order, _ := models.GetOneById(c.Param("id")) // method get by email
	response := new(models.Response)

	if order.DeleteOrder(c.Param("id")) != nil { // method update order
		response.ErrorCode = 10
		response.Message = "Gagal menghapus data order"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses menghapus data order"
	}
	return c.JSON(http.StatusOK, response)
}

func GetOrder(c echo.Context) error {
	response := new(models.Response)
	order, err := models.GetOrderAll() // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal melihat data orders"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses melihat data orders"
		response.Data = order
	}
	return c.JSON(http.StatusOK, response)
}
