package config

import (
	"os"
	"strconv"
)

type Config struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
	WeatherAPIKey string
	WeatherAPIurl string
	ServerPort    string
}

func LoadConfig() (*Config, error) {
	_ = os.Setenv("REDIS_HOST", "localhost")
	_ = os.Setenv("REDIS_PORT", "6379")
	_ = os.Setenv("REDIS_PASSWORD", "")
	_ = os.Setenv("REDIS_DB", "0")
	_ = os.Setenv("WEATHER_API_KEY", "{Your_OpenWeather_API_Key}")
	_ = os.Setenv("WEATHER_API_URL", "https://api.openweathermap.org/data/2.5/weather")

	_ = os.Setenv("SERVER_PORT", "8040")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return &Config{
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       redisDB,
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
		WeatherAPIurl: os.Getenv("WEATHER_API_URL"),
		ServerPort:    os.Getenv("SERVER_PORT"),
	}, nil
}
