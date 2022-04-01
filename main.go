package main

import (
	"thailephan/flashcard-echo-api/entities"
	"thailephan/flashcard-echo-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var users []entities.User

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitRoutes(e)
	
	e.Logger.Fatal(e.Start("127.0.0.1:8081"))
}