/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cityCmd represents the city command
var cityCmd = &cobra.Command{
	Use:   "city",
	Short: "Manage a list of your favourite cities in the world.",
	Long: `Manage a list your favourite cities in the world.

You can add, remove and list all your favourite city with subcommands below.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(cityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
