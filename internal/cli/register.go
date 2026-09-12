package cli

import (
	"fmt"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/bot"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register slash commands with discord",
	Run: func(_ *cobra.Command, _ []string) {
		token := viper.GetString("general.token")

		if len(token) == 0 {
			fmt.Println("Missing token in config, aborting.")
			return
		}

		bot, err := bot.New(token)

		if err != nil {
			log.Fatal("Bot couldn't initalize..")
		}

		guildID := viper.GetString("guildID")

		err = bot.RegisterCommands(guildID)

		if err != nil {
			log.Fatal("Couldn't register commands")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
