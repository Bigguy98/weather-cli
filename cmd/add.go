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
}
