package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherResponse struct {
	Weather  []Weather `json:"weather"`
	Main     Main      `json:"main"`
	TempType string    `json:"tempType"`
}

type Weather struct {
	Main string `json:"main"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

func main() {
	resp, err := getCurrentWeather("api_key", 38.6270, -90.1994)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current Weather:", resp.Weather[0].Main)
	fmt.Println("Temperature Type:", resp.TempType)
}

func getCurrentWeather(apiKey string, lat float64, long float64) (WeatherResponse, error) {
	url := "http://api.openweathermap.org/data/2.5/weather?lat=" + fmt.Sprintf("%f", lat) + "&lon=" + fmt.Sprintf("%f", long) + "&units=imperial&appid=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, err
	}

	var weatherResp WeatherResponse
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		return WeatherResponse{}, err
	}

	var currentWeather string
	if len(weatherResp.Weather) > 0 {
		currentWeather = weatherResp.Weather[0].Main
	}

	temperatureType := getTemperatureType(weatherResp.Main.Temp)

	weatherResp.Weather[0].Main = currentWeather
	weatherResp.TempType = temperatureType

	return weatherResp, nil
}

func getTemperatureType(temp float64) string {
	if temp < 45 {
		return "cold"
	} else if temp >= 45 && temp <= 65 {
		return "moderate"
	} else if temp > 65 && temp <= 80 {
		return "perfect!"
	} else {
		return "hot"
	}
}
