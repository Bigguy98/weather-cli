package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type City struct {
	Name       string
	Latitude   string
	Longitude  string
	Country    string
	Timezone   string
	Population string
}

var CITY_FILE = "cities.csv"

type ResponseCity struct {
	Results []struct {
		ID         int32   `json:"id"`
		Name       string  `json:"name"`
		Latitude   float32 `json:"latitude"`
		Longitude  float32 `json:"longitude"`
		Elevation  float32 `json:"elevation"`
		Country    string  `json:"country"`
		Timezone   string  `json:"timezone"`
		Population int64   `json:"population"`
	} `json:"results"`
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func AddCity(cityName string) {
	// check if city already exist
	city := getCityByName(cityName)
	if (City{}) != city {
		fmt.Printf("City %v already exist on list. Run 'weather-cli city list' to view!\n", cityName)
		return
	}

	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%v&count=1&language=en&format=json", cityName)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{Timeout: time.Second * 10}

	resp, _ := client.Do(req)

	body, _ := io.ReadAll(resp.Body)

	resp.Body.Close()

	var responseCity ResponseCity
	if err := json.Unmarshal(body, &responseCity); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON for response city")
	} else {
		// fmt.Printf("%v\n", prettyPrint(responseCity))
		if len(responseCity.Results) < 1 {
			fmt.Printf("No city with name %v found!\n", cityName)
		} else {
			city := City{Name: cityName,
				Latitude:   fmt.Sprintf("%v", fmt.Sprintf("%.5f", responseCity.Results[0].Latitude)),
				Longitude:  fmt.Sprintf("%v", fmt.Sprintf("%.5f", responseCity.Results[0].Longitude)),
				Country:    fmt.Sprintf("%v", responseCity.Results[0].Country),
				Timezone:   fmt.Sprintf("%v", responseCity.Results[0].Timezone),
				Population: fmt.Sprintf("%v", responseCity.Results[0].Population),
			}
			saveToFile(city)
			fmt.Printf("Adding city %v success!\n", cityName)
		}

	}
}

func saveToFile(city City) {
	file, _ := os.OpenFile(CITY_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := csv.NewWriter(file)
	row := []string{city.Name, city.Latitude, city.Longitude, city.Country, city.Timezone, city.Population}
	writer.Write(row)
	writer.Flush()
	file.Close()
}

func getCities() []City {
	file, _ := os.OpenFile(CITY_FILE, os.O_RDONLY, 0644)
	reader := csv.NewReader(file)
	cities := []City{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		cities = append(cities, City{
			Name: row[0], Latitude: row[1], Longitude: row[2], Country: row[3], Timezone: row[4], Population: row[5],
		})
	}
	file.Close()
	return cities
}

func getCityByName(cityName string) City {
	file, _ := os.OpenFile(CITY_FILE, os.O_RDONLY, 0644)
	reader := csv.NewReader(file)
	var city City = City{}
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if row[0] == cityName {
			city = City{
				Name: row[0], Latitude: row[1], Longitude: row[2], Country: row[3], Timezone: row[4], Population: row[5],
			}
			break
		}
	}
	file.Close()
	return city
}

func RemoveCity(cityName string) {
	cities := getCities()

	file, _ := os.OpenFile(CITY_FILE, os.O_RDWR|os.O_TRUNC, 0644)
	writer := csv.NewWriter(file)
	for _, city := range cities {
		if city.Name != cityName {
			row := []string{city.Name, city.Latitude, city.Longitude, city.Country, city.Timezone, city.Population}
			writer.Write(row)
		}
	}
	writer.Flush()
	file.Close()
}

func ListCity() {
	cities := getCities()
	fmt.Println("Name\t Latitude\t Longitude\t Population\t Country\t ")
	for _, city := range cities {
		fmt.Printf("%s\t %s\t %s\t %s\t %s\n", city.Name, city.Latitude, city.Longitude, city.Population, city.Country)
	}
}
