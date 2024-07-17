package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"simple-weather-api/config"
)

type GeocodeResponse struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

func GetLatLongFromCity(city string) (float64, float64, error) {
	url := fmt.Sprintf("%s/geo/1.0/direct?q=%s&limit=2&appid=%s", config.OpenWeatherMapBaseUrl, city, os.Getenv("OPEN_WEATHER_API_KEY"))
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read response body: %w", err)
	}

	var geocodeResponses []GeocodeResponse
	if err := json.Unmarshal(body, &geocodeResponses); err != nil {
		return 0, 0, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	if len(geocodeResponses) == 0 {
		return 0, 0, fmt.Errorf("no coordinates found for city: %s", city)
	}

	return geocodeResponses[0].Lat, geocodeResponses[0].Lon, nil
}
