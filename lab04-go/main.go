package main

import (
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
)

type WeatherResponse struct {
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
}

func getWeather(c echo.Context) error {
	response := WeatherResponse{
		Location:    "Warsaw",
		Temperature: 18.5,
		Condition:   "Cloudy",
	}

	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()

	e.GET("/weather", getWeather)
	e.POST("/weather", getWeather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}