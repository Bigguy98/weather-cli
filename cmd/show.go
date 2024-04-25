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
	Short: "A brief description of your command",
	Long: `Show detail weather for cities you have save:
	
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
