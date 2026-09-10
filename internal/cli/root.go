// Package cli provides all cli commands to the binary.
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nono-bot",
	Short: "Nono bot is a discord bot written in go",
	Long:  "Nono bot is a go discord bot designed to keep nono words out of your server.",
}

// Execute the cli program.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
