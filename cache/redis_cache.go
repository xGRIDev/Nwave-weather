package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xGRIDEv/NWave-Weather/models"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
	ttl    time.Duration
}

func NewRedisCache(host, port, password string, db int) (*RedisCache, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	//test-connect
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("gagal konek pada Redis: %v", err)
	}

	return &RedisCache{
		client: client,
		ctx:    ctx,
		ttl:    10 * time.Minute,
	}, nil
}

func (r *RedisCache) GetWeather(city string) (*models.WeatherCache, error) {
	key := fmt.Sprintf("weather:%s", city)

	data, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error getting cache: %v", err)
	}
	var cachedWeather models.WeatherCache
	err = json.Unmarshal([]byte(data), &cachedWeather)
	if err != nil {
		return nil, fmt.Errorf("error marshalling cache: %v", err)
	}

	//check cache expired
	if time.Now().After(cachedWeather.Exp) {
		r.DeleteWeather(city)
		return nil, nil
	}

	return &cachedWeather, nil
}

func (r *RedisCache) SetWeather(city string, weather *models.WeatherResponse) error {
	key := fmt.Sprintf("weather:%s", city)
	cachedWeather := models.WeatherCache{
		Data:      weather,
		TimeStamp: time.Now(),
		Exp:       time.Now().Add(r.ttl),
	}

	data, err := json.Marshal(cachedWeather)
	if err != nil {
		return fmt.Errorf("error masrhasling cache: %v", err)
	}
	err = r.client.Set(r.ctx, key, data, r.ttl).Err()
	if err != nil {
		return fmt.Errorf("error masrhasling cache: %v", err)

	}
	return nil
}

func (r *RedisCache) DeleteWeather(city string) error {
	key := fmt.Sprintf("weather:%s", city)
	return r.client.Del(r.ctx, key).Err()
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}
