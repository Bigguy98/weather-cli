/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	service "weather-cli/cmd/service"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove city from your favourite list",
	Long: `Syntax to use 'weather-cli city remove [city name]'. For example:

weather-cli city remove Hanoi`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			service.RemoveCity(args[0])
		} else {
			fmt.Printf("Expected 1 argument but found 0!")
		}

	},
}

func init() {
	cityCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
