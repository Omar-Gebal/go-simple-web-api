package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"simple-weather-api/config"
)

type WeatherData struct {
	WeatherDescription string  `json:"weather_description"`
	Temperature        float64 `json:"temperature"`
}

func GetWeatherData(lat, lon float64) (WeatherData, error) {
	url := fmt.Sprintf("%s/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric", config.OpenWeatherMapBaseUrl, lat, lon, os.Getenv("OPEN_WEATHER_API_KEY"))
	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherData{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var weatherData map[string]interface{}
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return WeatherData{}, fmt.Errorf("failed to parse weather data: %w", err)
	}

	weatherDesc := "Unknown"
	if weatherArray, ok := weatherData["weather"].([]interface{}); ok && len(weatherArray) > 0 {
		if weatherMap, ok := weatherArray[0].(map[string]interface{}); ok {
			if desc, ok := weatherMap["description"].(string); ok {
				weatherDesc = desc
			}
		}
	}

	temperature := 0.0
	if mainData, ok := weatherData["main"].(map[string]interface{}); ok {
		if temp, ok := mainData["temp"].(float64); ok {
			temperature = temp
		}
	}

	return WeatherData{
		WeatherDescription: weatherDesc,
		Temperature:        temperature,
	}, nil
}
