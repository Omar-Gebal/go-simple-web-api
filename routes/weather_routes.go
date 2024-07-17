package routes

import (
	"simple-weather-api/handlers"
	"simple-weather-api/middleware"

	"github.com/labstack/echo/v4"
)

func WeatherRoutes(e *echo.Echo) {
	e.GET("/weather", handlers.HandleGetWeather, middleware.AuthMiddleware)
}
