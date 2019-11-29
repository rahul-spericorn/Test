package main

import (
	"loosidAPI/config"
	"loosidAPI/db"
	"loosidAPI/generated"
	"loosidAPI/handlers"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	config.SetConfig()
	db.InitDb()
	defer db.DbConnection.Close()

	e := echo.New()
	router := e.Group("/v3")

	handler := handlers.ServerWrapper{}

	generated.RegisterHandlers(router, handler)

	// Start server
	e.Logger.Fatal(e.Start(":" + config.Cfg.Port))
}
