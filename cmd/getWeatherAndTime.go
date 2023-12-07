/*
Copyright © 2023 Daniel daniel-vuky@gmail.com
*/
package cmd

import (
	weatherAndTime "github.com/daniel-vuky/weather-and-time-golang-cli/provider"
	"github.com/spf13/cobra"
)

// getWeatherAndTimeCmd represents the getWeatherAndTime command
var getWeatherAndTimeCmd = &cobra.Command{
	Use:   "weather-and-time",
	Short: "Get the weather and time",
	Long:  `Pass US Zipcode, UK Postcode, Canada Postalcode, IP address, Latitude/Longitude (decimal degree) or city name`,
	Run: func(cmd *cobra.Command, args []string) {
		weatherAndTime.GetWeatherAndTime(args[0])
	},
}

func init() {
	rootCmd.AddCommand(getWeatherAndTimeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getWeatherAndTimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getWeatherAndTimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
