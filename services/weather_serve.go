package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/xGRIDEv/NWave-Weather/cache"
	"github.com/xGRIDEv/NWave-Weather/models"
)

type WeatherService struct {
	cache      *cache.RedisCache
	apiKey     string
	apiURL     string
	httpClient *http.Client
}

func NewWeatherService(cache *cache.RedisCache, apiKey, apiURL string) *WeatherService {
	return &WeatherService{
		cache:      cache,
		apiKey:     apiKey,
		apiURL:     apiURL,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *WeatherService) GetWeather(city string) (*models.WeatherResponse, error) {
	cached, err := s.cache.GetWeather(city)
	if err != nil {
		return nil, fmt.Errorf("cache error: %v", err)
	}
	if cached != nil {
		fmt.Printf("Cached HIT for city : %s\n", city)
		return cached.Data, nil
	}

	fmt.Printf("Cache Miss for City : %s, fetching data from API...\n", city)

	weather, err := s.fetFromAPI(city)
	if err != nil {
		return nil, fmt.Errorf("API error: %v", err)
	}
	err = s.cache.SetWeather(city, weather)
	if err != nil {
		fmt.Printf("Warning: failed to cache  weather for %s: %v\n", city, err)
	}

	return weather, nil
}
func (s *WeatherService) fetFromAPI(city string) (*models.WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=44.34&lon=10.99&appid=%s", s.apiKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	rspn, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer rspn.Body.Close()
	if rspn.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", rspn.StatusCode)
	}

	body, err := io.ReadAll(rspn.Body)
	if err != nil {
		return nil, err
	}
	var weather models.WeatherResponse
	err = json.Unmarshal(body, &weather)

	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func (s *WeatherService) ClearCache(city string) error {
	return s.cache.DeleteWeather(city)
}
