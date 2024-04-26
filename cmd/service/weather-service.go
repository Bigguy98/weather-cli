package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// sharing this array for go routines
var weathers []ResponseWeather
var mutex sync.Mutex

type ResponseWeather struct {
	CityName string
	Current  struct {
		Temperature float32 `json:"temperature_2m"` // temperature above 2m
		WindSpeed   float32 `json:"wind_speed_10m"`
		Humidity    int16   `json:"relative_humidity_2m"`
		Visibility  float32 `json:"visibility"`
		Rain        float32 `json:"rain"`
	} `json:"current"`
}

type ResponseForecast struct {
	CityName string
	Hourly   struct {
		Time        []string  `json:"time"`
		Temperature []float64 `json:"temperature_2m"` // temperature above 2m
		WindSpeed   []float64 `json:"wind_speed_10m"`
		Humidity    []int     `json:"relative_humidity_2m"`
	} `json:"hourly"`
}

func ForecastWeather(cityName string) (ResponseForecast, error) {
	var forecast = ResponseForecast{CityName: cityName}

	city := getCityByName(cityName)
	if (City{}) == city { // check if city is empty
		fmt.Println("City not found on list, please add city with command 'weather-cli city add [city name]'")
		return forecast, errors.New("City not found on list")
	} else {
		url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&hourly=temperature_2m,wind_speed_10m,relative_humidity_2m,visibility,rain", city.Latitude, city.Longitude)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Cache-Control", "no-cache")

		client := &http.Client{Timeout: time.Second * 10}

		resp, _ := client.Do(req)

		body, _ := io.ReadAll(resp.Body)

		resp.Body.Close()

		if err := json.Unmarshal(body, &forecast); err != nil { // Parse []byte to the go struct pointer
			fmt.Printf("Can not unmarshal JSON for forecast weather casuse by %s\n", err)
		}
	}
	return forecast, nil

}

func getWeatherDetails(city City) ResponseWeather {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&current=temperature_2m,wind_speed_10m,relative_humidity_2m,visibility,rain", city.Latitude, city.Longitude)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{Timeout: time.Second * 10}

	resp, _ := client.Do(req)

	body, _ := io.ReadAll(resp.Body)

	resp.Body.Close()

	var responseWeather = ResponseWeather{CityName: city.Name}
	if err := json.Unmarshal(body, &responseWeather); err != nil { // Parse []byte to the go struct pointer
		fmt.Printf("Can not unmarshal JSON for repsonse weather casuse by %s\n", err)
	}

	return responseWeather

}

func gatherWeatheForCities(city City, wg *sync.WaitGroup) {
	defer wg.Done()
	weather := getWeatherDetails(city)
	// implement mutex here to avoid conflict between routines
	mutex.Lock()
	weathers = append(weathers, weather)
	defer mutex.Unlock()
}

func ShowWeatherForAllCity() {
	cities := getCities()

	weathers = []ResponseWeather{}

	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go gatherWeatheForCities(city, &wg)
	}
	// wait until all routines finish
	wg.Wait()

	fmt.Println("City\t Temp(Celcius)\t Wind(Km/h)\t Humidity(%)\t Visibility(Km)\t Rain(mm)\t")
	for _, cityWeather := range weathers {
		fmt.Printf("%s\t %s\t\t %s\t\t %s\t\t %s\t\t %s\t\n",
			cityWeather.CityName,
			fmt.Sprintf("%.1f", cityWeather.Current.Temperature),
			fmt.Sprintf("%.1f", cityWeather.Current.WindSpeed),
			fmt.Sprintf("%d", cityWeather.Current.Humidity),
			fmt.Sprintf("%.0f", cityWeather.Current.Visibility),
			fmt.Sprintf("%.0f", cityWeather.Current.Rain))
	}

}

func ShowWeatherForCity(cityName string) {
	city := getCityByName(cityName)

	if (City{}) == city { // check if city is empty
		fmt.Println("City not found on list, please add city with command 'weather-cli city add [city name]'")
	} else {
		cityWeather := getWeatherDetails(city)
		fmt.Println("City\t Temp(Celcius)\t WindSpeed(Km/h)\t Humidity(%)\t  Visibility(Km)\t  Rain(mm)\t")
		fmt.Printf("%s\t\t %s\t\t %s\t\t %s\t\t %s\t\t %s\t\t\n",
			cityName,
			fmt.Sprintf("%.1f", cityWeather.Current.Temperature),
			fmt.Sprintf("%.1f", cityWeather.Current.WindSpeed),
			fmt.Sprintf("%d", cityWeather.Current.Humidity),
			fmt.Sprintf("%.0f", cityWeather.Current.Visibility),
			fmt.Sprintf("%.0f", cityWeather.Current.Rain))
	}
}
