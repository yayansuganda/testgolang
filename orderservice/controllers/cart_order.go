package controllers

import (
	"net/http"
	"orderservice/models"

	"github.com/labstack/echo"
)

func StoreCartOrder(c echo.Context) error {
	cart_order := new(models.Cart_order)
	c.Bind(cart_order)
	response := new(models.Response)
	if cart_order.CreateCartOrder() != nil { // method create cart order
		response.ErrorCode = 10
		response.Message = "Gagal create data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses create data order"
		response.Data = *cart_order
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateCartOrder(c echo.Context) error {
	cart_order := new(models.Cart_order)
	c.Bind(cart_order)
	response := new(models.Response)
	if cart_order.UpdateCartOrder(c.Param("id")) != nil { // method update cart order
		response.ErrorCode = 10
		response.Message = "Gagal update data order"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses update data order"
		response.Data = *cart_order
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteCartOrder(c echo.Context) error {
	cart_order := new(models.Cart_order)
	// order, _ := models.GetOneById(c.Param("id")) // method get by email
	response := new(models.Response)

	if cart_order.DeleteCartOrder(c.Param("id")) != nil { // method update order
		response.ErrorCode = 10
		response.Message = "Gagal menghapus data order"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses menghapus data order"
	}
	return c.JSON(http.StatusOK, response)
}

func GetCartOrder(c echo.Context) error {
	response := new(models.Response)
	order, err := models.GetAll() // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal melihat data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses melihat data user"
		response.Data = order
	}
	return c.JSON(http.StatusOK, response)
}
