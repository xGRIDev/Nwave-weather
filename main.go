package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	routes "github.com/xGRIDEv/NWave-Weather/Routes"
	"github.com/xGRIDEv/NWave-Weather/cache"
	"github.com/xGRIDEv/NWave-Weather/config"
	"github.com/xGRIDEv/NWave-Weather/handler"
	"github.com/xGRIDEv/NWave-Weather/services"
)

func main() {
	cnfr, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	redisCache, err := cache.NewRedisCache(
		cnfr.RedisHost,
		cnfr.RedisPort,
		cnfr.RedisPassword,
		cnfr.RedisDB,
	)
	if err != nil {
		log.Fatalf("Redis cache initialized succesfully: %v", err)
	}
	defer redisCache.Close()
	fmt.Println("Redis cache intialized successfully")

	weatherService := services.NewWeatherService(
		redisCache,
		cnfr.WeatherAPIKey,
		cnfr.WeatherAPIurl,
	)

	//Handler Init
	weatherHandler := handler.NewWeatherHndlr(weatherService)

	//Setup Gin-Reouter
	router := gin.Default()

	//Setup routes
	routes.SetRoutes(router, weatherHandler)

	//Running-server
	fmt.Printf("server Starting on Port : %s\n", cnfr.ServerPort)
	if err := router.Run(":" + cnfr.ServerPort); err != nil {
		log.Fatalf("failed to run server : %v", err)
	}
}
