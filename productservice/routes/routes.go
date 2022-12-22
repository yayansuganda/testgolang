package routes

import (
	"productservice/controllers"
	"productservice/middleware"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/product", controllers.GetProduct, middleware.IsAuthenticated)
	e.GET("/product/detail/:id", controllers.GetSingleProduct, middleware.IsAuthenticated)
	e.POST("/product", controllers.StoreProduct, middleware.IsAuthenticated)
	e.PUT("/product/:id", controllers.UpdateProduct, middleware.IsAuthenticated)
	e.DELETE("/product/:id", controllers.DeleteProduct, middleware.IsAuthenticated)

	return e
}
