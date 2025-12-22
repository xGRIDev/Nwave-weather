package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xGRIDEv/NWave-Weather/handler"
)

func SetRoutes(router *gin.Engine, hdl *handler.WeatherHandler) {
	api := router.Group("/api/v1")
	{
		api.GET("/weather", hdl.GetWeather)
		api.GET("/health", hdl.HealthCheck)
		api.DELETE("/weather/cache", hdl.ClearCache)
	}
}
