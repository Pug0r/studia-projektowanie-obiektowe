package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/glebarez/sqlite"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Weather struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
}

type WeatherProxy struct {
	client *http.Client
}

type openMeteoResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WeatherCode int     `json:"weathercode"`
	} `json:"current_weather"`
}

type locationCoordinates struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type weatherRequest struct {
	Locations []string `json:"locations"`
}

var supportedLocations = map[string]locationCoordinates{
	"warsaw": {Name: "Warsaw", Latitude: 52.23, Longitude: 21.01},
	"krakow": {Name: "Krakow", Latitude: 50.06, Longitude: 19.94},
	"gdansk": {Name: "Gdansk", Latitude: 54.35, Longitude: 18.65},
}

func NewWeatherProxy() *WeatherProxy {
	return &WeatherProxy{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (p *WeatherProxy) GetWeather(location string, latitude, longitude float64) (Weather, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.2f&longitude=%.2f&current_weather=true", latitude, longitude)
	resp, err := p.client.Get(url)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("external api returned status %d", resp.StatusCode)
	}

	var payload openMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return Weather{}, err
	}

	return Weather{
		Location:    location,
		Temperature: payload.CurrentWeather.Temperature,
		Condition:   fmt.Sprintf("code-%d", payload.CurrentWeather.WeatherCode),
	}, nil
}

var db *gorm.DB
var weatherProxy *WeatherProxy

func initDatabase() error {
	database, err := gorm.Open(sqlite.Open("weather.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := database.AutoMigrate(&Weather{}); err != nil {
		return err
	}

	seedData := []Weather{
		{Location: "Warsaw", Temperature: 18.5, Condition: "Cloudy"},
		{Location: "Krakow", Temperature: 20.0, Condition: "Sunny"},
		{Location: "Gdansk", Temperature: 16.0, Condition: "Windy"},
	}

	if err := database.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Weather{}).Error; err != nil {
		return err
	}

	if err := database.Create(&seedData).Error; err != nil {
		return err
	}

	db = database
	return nil
}

func getWeather(c echo.Context) error {
	locations := []string{"warsaw", "krakow", "gdansk"}

	queryLocations := c.QueryParam("locations")
	if queryLocations != "" {
		locations = strings.Split(queryLocations, ",")
	}

	if c.Request().Method == http.MethodPost && c.Request().ContentLength > 0 {
		var req weatherRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}
		if len(req.Locations) > 0 {
			locations = req.Locations
		}
	}

	result := make([]Weather, 0, len(locations))
	for _, rawLocation := range locations {
		key := strings.ToLower(strings.TrimSpace(rawLocation))
		coordinates, ok := supportedLocations[key]
		if !ok {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "unsupported location"})
		}

		weather, err := weatherProxy.GetWeather(coordinates.Name, coordinates.Latitude, coordinates.Longitude)
		if err != nil {
			return c.JSON(http.StatusBadGateway, map[string]string{"error": "external api error"})
		}

		if err := db.Create(&weather).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
		}

		result = append(result, weather)
	}

	return c.JSON(http.StatusOK, result)
}

func main() {
	if err := initDatabase(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	weatherProxy = NewWeatherProxy()

	e.GET("/weather", getWeather)
	e.POST("/weather", getWeather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}