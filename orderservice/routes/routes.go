package routes

import (
	"orderservice/controllers"
	"orderservice/middleware"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/cart_order", controllers.GetCartOrder, middleware.IsAuthenticated)
	e.POST("/cart_order", controllers.StoreCartOrder, middleware.IsAuthenticated)
	e.PUT("/cart_order/:id", controllers.UpdateCartOrder, middleware.IsAuthenticated)
	e.DELETE("/cart_order/:id", controllers.DeleteCartOrder, middleware.IsAuthenticated)

	e.GET("/order", controllers.GetOrder, middleware.IsAuthenticated)
	e.POST("/order", controllers.StoreOrder, middleware.IsAuthenticated)
	e.PUT("/order/:id", controllers.UpdateOrder, middleware.IsAuthenticated)
	e.DELETE("/order/:id", controllers.DeleteOrder, middleware.IsAuthenticated)

	return e
}
