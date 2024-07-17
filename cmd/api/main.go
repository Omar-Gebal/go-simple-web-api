package main

import (
	"log"
	"net/http"
	"simple-weather-api/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Wassup!")
	})

	routes.WeatherRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))

}
