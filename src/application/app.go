package application

import (
	"delivery-app/src/application/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func BootstrapAPP() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	routes.Router(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
