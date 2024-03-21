package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"modul/controller"
)

func main() {
	e := echo.New()

	e.GET("/games", controller.GetGames)
	e.GET("/games/:id", controller.GetGame)
	e.POST("/games/:name/:genre/:price", controller.CreateGame)
	e.PUT("/games/:id/:name/:genre/:price", controller.UpdateGame)
	e.DELETE("/games/:id", controller.DeleteGame)

	e.Logger.Fatal(e.Start(":8080"))
}
