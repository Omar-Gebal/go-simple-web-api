package handlers

import (
	"log"
	"net/http"
	"simple-weather-api/services"

	"github.com/labstack/echo/v4"
)

type GetWeatherResponse struct {
	WeatherDescription string  `json:"weather_description"`
	Temperature        float64 `json:"temperature"`
	Msg                string  `json:"msg"`
}

func HandleGetWeather(c echo.Context) error {
	city := c.QueryParam("city")
	//TODO: separate logic
	//is it a good idea to make a separate service just to combine 3 lines ?
	lat, lon, err := services.GetLatLongFromCity(city)
	if err != nil {
		log.Printf("Error getting coordinates: %v", err)
		return c.JSON(http.StatusInternalServerError, GetWeatherResponse{WeatherDescription: "", Msg: "Failed to get coordinates"})
	}

	weatherData, err := services.GetWeatherData(lat, lon)
	if err != nil {
		log.Printf("Error getting weather data: %v", err)
		return c.JSON(http.StatusInternalServerError, GetWeatherResponse{WeatherDescription: "", Msg: "Failed to get weather data"})
	}

	return c.JSON(http.StatusOK, GetWeatherResponse{WeatherDescription: weatherData.WeatherDescription, Temperature: weatherData.Temperature, Msg: "success"})
}
