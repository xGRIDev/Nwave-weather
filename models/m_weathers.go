package models

import "time"

type WeatherResponse struct {
	Coord struct {
		Long  float64 `json:"long"`
		Latit float64 `json:"latit"`
	}
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:description"`
		Icon        string `json:"icon"`
	}
	Base string
	Main struct {
		Temp      float64 `json:"temp"`
		Feelslike float64 `json:"feels_like"`
		MinTemp   float64 `json:"min_temp"`
		MaxTemp   float64 `json:"max_temp"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	}
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"id"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type WeatherCache struct {
	Data      *WeatherResponse `json:"data"`
	TimeStamp time.Time        `json:"timestamp"`
	Exp       time.Time        `json:"expired_at"`
}
