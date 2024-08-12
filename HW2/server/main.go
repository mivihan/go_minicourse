package main

import (
	"awesomeProject/accounts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/google/uuid"
	"go_minicourse/HW2/server/account"
)

var secretKey = uuid.New().String()

func main() {
	println("Секретный ключ:", secretKey)
	
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/account/create", accountsHandler.CreateAccount)
	e.GET("/api/account", accountsHandler.GetAccount)
	e.DELETE("/api/account", accountsHandler.DeleteAccount)
	e.PATCH("/api/account", accountsHandler.PatchAccount)
	e.POST("/api/account/rename", accountsHandler.ChangeAccount)

	e.GET("/api/accounts", accountsHandler.GetAll)
	e.GET("/api/actuator", accountsHandler.Actuator)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
