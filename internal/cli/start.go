package cli

import (
	"fmt"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/bot"
	"github.com/MsZoezo/nono-bot/internal/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start [token]",
	Short: "Start nono bot",
	Run: func(_ *cobra.Command, _ []string) {
		token := viper.GetString("general.token")

		if len(token) == 0 {
			fmt.Println("Missing token in config, aborting.")
			return
		}

		db, err := db.New()

		if err != nil {
			return
		}

		bot, err := bot.New(db, token)

		if err != nil {
			log.Fatal("Bot couldn't initalize..")
		}

		err = bot.Run()

		if err != nil {
			log.Fatal("Error while running bot..")
		}

		bot.Close()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
