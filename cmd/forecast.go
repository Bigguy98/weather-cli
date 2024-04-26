/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"weather-cli/cmd/service"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
)

var CHART_DATA_SIZE = 50

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Forecast hourly weather for a city",
	Long: `Forecast hourly weather (including temperature, wind speed and humidity) for a city. For example:

weather-cli forecast Hanoi`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cityName := args[0]
			forecastResult, error := service.ForecastWeather(cityName)
			if error == nil {
				showForecastChart(forecastResult)
			}
		} else {
			fmt.Println("Expected 1 argument, found 0!")
		}

	},
}

func showCharts(avgHumidityPercent int, hourlyTemper []float64, hourlyWindSpeed []float64, cityName string) {
	if err := ui.Init(); err != nil {
		fmt.Printf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	parag := widgets.NewParagraph()
	parag.Text = fmt.Sprintf("Forecast hourly weather for %v (Press Q to exit)", cityName)
	parag.SetRect(20, 0, 100, 1)
	parag.Border = false

	g1 := widgets.NewGauge()
	g1.Title = "Humidity (%)"
	g1.SetRect(0, 1, 100, 4)
	g1.Percent = avgHumidityPercent
	g1.BarColor = ui.ColorGreen
	g1.LabelStyle = ui.NewStyle(ui.ColorYellow)
	g1.TitleStyle.Fg = ui.ColorMagenta
	g1.BorderStyle.Fg = ui.ColorWhite

	p0 := widgets.NewPlot()
	p0.Title = "Hourly temperature (Celcius)"
	p0.TitleStyle = ui.NewStyle(ui.ColorYellow)
	p0.Data = [][]float64{hourlyTemper}
	p0.SetRect(0, 4, 50, 20)
	p0.AxesColor = ui.ColorWhite

	p1 := widgets.NewPlot()
	p1.Title = "Hourly windspeed (km/h)"
	p1.TitleStyle = ui.NewStyle(ui.ColorBlue)
	p1.Marker = widgets.MarkerDot
	p1.Data = [][]float64{hourlyWindSpeed}
	p1.SetRect(50, 4, 100, 20)
	p1.AxesColor = ui.ColorWhite
	p1.LineColors[0] = ui.ColorRed
	p1.DrawDirection = widgets.DrawLeft

	ui.Render(parag, g1, p0, p1)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func getAverage(array []int) int {
	length := len(array)
	sum := 0
	for _, num := range array {
		sum = sum + num
	}
	return sum / length
}

func showForecastChart(forecastData service.ResponseForecast) {
	showCharts(getAverage(forecastData.Hourly.Humidity), forecastData.Hourly.Temperature, forecastData.Hourly.WindSpeed, forecastData.CityName)
}

func init() {
	rootCmd.AddCommand(forecastCmd)
}
