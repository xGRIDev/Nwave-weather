package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xGRIDEv/NWave-Weather/services"
)

type WeatherHandler struct {
	WeatherService *services.WeatherService
}

func NewWeatherHndlr(weatherService *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		WeatherService: weatherService,
	}
}

func (hdl *WeatherHandler) GetWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city parameter is required"})
		return
	}
	weather, err := hdl.WeatherService.GetWeather(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"city":    city,
		"weather": weather,
	})
}

func (hdl *WeatherHandler) ClearCache(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city parameter is required"})
		return
	}
	err := hdl.WeatherService.ClearCache(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Cache cleared for city: %s", city),
	})
}

func (hdl *WeatherHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"cache":  "redis",
	})
}
