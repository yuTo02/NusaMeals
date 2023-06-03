package main

import (
	"reglog/config"
	"reglog/lib/seeder"
	"reglog/middlewares"
	route "reglog/routes"
	"reglog/util"

	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.InitDB()
	config.InitialMigration()
	seeder.DBSeed(config.DB)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Custom Middleware
	e.Use(middlewares.CORS())

	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}

	// Routes
	route.New()

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
