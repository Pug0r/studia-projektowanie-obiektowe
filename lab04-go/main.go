package main

import (
	"log"
	"net/http"
	"os"

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

var db *gorm.DB

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
	var weather []Weather
	if err := db.Find(&weather).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	return c.JSON(http.StatusOK, weather)
}

func main() {
	if err := initDatabase(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.GET("/weather", getWeather)
	e.POST("/weather", getWeather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}