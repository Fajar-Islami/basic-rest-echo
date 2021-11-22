package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Customer struct {
	Nama   string `validate:"required"`
	Email  string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur   int    `validate:"gte=17,lte=35"`
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Nama:   "Bams",
		Email:  "bam@mail.com",
		Alamat: "Jakart",
		Umur:   25,
	}

	// Validasi struct dengan validate
	err := v.Struct(cust)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"messages": "Success",
	})
}

func TestVariabelValidation(c echo.Context) error {
	v := validator.New()

	email := "bams@mail.com"

	// Validasi single variable
	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messages": err.Error(),
			"ket":      "Email not valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"messages": "Success",
	})
}
