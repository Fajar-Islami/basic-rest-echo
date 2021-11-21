package controllers

import (
	"net/http"

	"github.com/basic-echo-golang/models"
	"github.com/labstack/echo/v4"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {

	// Tampung parameter dari aplikasi eksternal
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.StorePegawai(nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// return dari modal
	return c.JSON(http.StatusOK, result)
}
