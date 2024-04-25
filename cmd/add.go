/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	service "weather-cli/cmd/service"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add city to your favourite list",
	Long: `Syntax to use 'weather-cli city add [city name]'. For example:

weather-cli city add Hanoi`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			city := args[0]
			service.AddCity(city)

		} else {
			fmt.Println("Expected 1 argument, found 0!")
		}

	},
}

func init() {
	cityCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
