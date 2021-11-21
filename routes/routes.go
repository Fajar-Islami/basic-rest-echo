package routes

import (
	"net/http"

	"github.com/basic-echo-golang/controllers"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello this is echo!")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)
	e.POST("pegawai", controllers.StorePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)

	return e
}
