package routes

import (
	"autentication/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	return e
}
