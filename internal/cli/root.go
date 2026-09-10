// Package cli provides all cli commands to the binary.
package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFileNotFoundError viper.ConfigFileNotFoundError

var rootCmd = &cobra.Command{
	Use:   "nono-bot",
	Short: "Nono bot is a discord bot written in go",
	Long:  "Nono bot is a go discord bot designed to keep nono words out of your server.",
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		viper.SetConfigName("config")

		viper.AddConfigPath("$home/.config/nono-bot/")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()

		if err != nil {
			if errors.As(err, &configFileNotFoundError) {
				fmt.Println("Config file not found, consider creating one.")
			} else {
				fmt.Println("Error reading config file, aborting")
			}

			os.Exit(1)
		}
	},
}

// Execute the cli program.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
