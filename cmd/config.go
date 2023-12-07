/*
Copyright © 2023 Daniel daniel-vuky@gmail.com
*/
package cmd

import (
	"anhvdk/weather-and-time-cli/config"
	"github.com/spf13/cobra"
)

var uri, apiKey string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Use this command to config settings",
	Long: `Available settings:
- Uri
- Api Key`,
	Run: func(cmd *cobra.Command, args []string) {
		config.AddSettings(uri, apiKey)
	},
}

func init() {

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configCmd.Flags().StringVarP(&uri, "uri", "u", "", "Uri")
	configCmd.MarkFlagRequired("uri")
	configCmd.Flags().StringVarP(&apiKey, "api_key", "a", "", "Api Key")
	configCmd.MarkFlagRequired("api_key")
	rootCmd.AddCommand(configCmd)
}
