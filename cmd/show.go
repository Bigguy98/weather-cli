/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"weather-cli/cmd/service"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current detail weather for city/cities you like",
	Long: `Show current detail weather for cities you have like:

You can use 'weather-cli show' to get weather for all cities
Or use 'weather-cli show [city name] to get weather of a city you choose'
To list all saved city, run command 'weather-cli city list'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cityName := args[0]
			service.ShowWeatherForCity(cityName)
		} else {
			service.ShowWeatherForAllCity()
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

}
