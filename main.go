package main

import (
	"cep-consult/api/config/middleware"
	"cep-consult/api/handle"
	"fmt"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title CEP Consult API
// @description This is a simple API for address lookup based on CEP
// @version 1
// @host localhost:8080
// @BasePath /v1
func main() {
	e := echo.New()

	// Middleware CORS
	e.Use(middleware.CORS())

	// Swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// ConsultAddress Route and documentation
	// @Summary Consult a CEP and return its address
	// @Description Consult a CEP and return its address
	// @ID consult-address
	// @Accept  json
	// @Produce  json
	// @Param cep body string true "CEP to consult"
	// @Success 200 {object} AddressResponse
	// @Failure 400 {string} string "Invalid Request"
	// @Failure 404 {string} string "CEP not found"
	// @Router /consult-address [post]
	e.POST("/v1/consult-address", handle.ConsultAddress)

	fmt.Println("Running Port:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
