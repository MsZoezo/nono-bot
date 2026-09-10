// Package cli provides all cli commands to the binary.
package cli

import (
	"errors"

	"charm.land/log/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFileNotFoundError viper.ConfigFileNotFoundError

var rootCmd = &cobra.Command{
	Use:   "nono-bot",
	Short: "Nono bot is a discord bot written in go",
	Long:  "Nono bot is a go discord bot designed to keep nono words out of your server.",
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		log.Default().SetLevel(log.DebugLevel)

		viper.SetConfigName("config")

		viper.AddConfigPath("$home/.config/nono-bot/")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()

		if err != nil {
			if errors.As(err, &configFileNotFoundError) {
				log.Fatal("Config file not found, consider creating one.")
			}

			log.Fatal("Error reading config file..")
		}

		log.Debug("Succesfully read config file.")
	},
}

// Execute the cli program.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
