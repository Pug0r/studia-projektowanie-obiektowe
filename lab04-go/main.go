package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func NewWeatherProxy() *WeatherProxy {
	return &WeatherProxy{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (p *WeatherProxy) GetWeather() (Weather, error) {
	url := "https://api.open-meteo.com/v1/forecast?latitude=52.23&longitude=21.01&current_weather=true"
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
		Location:    "Warsaw",
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
	weather, err := weatherProxy.GetWeather()
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "external api error"})
	}

	return c.JSON(http.StatusOK, weather)
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