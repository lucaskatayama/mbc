/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var instrument string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get public data",
}

func init() {
	cobra.OnInitialize(func() {
		if instrument == "" {
			fmt.Println("Instrument symbol should not be empty")
		}
	})

	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.PersistentFlags().StringVarP(&instrument, "instrument", "i", "", "Instrument symbol in the format BASE-QUOTE (e.g. BTC-BRL)")
}
