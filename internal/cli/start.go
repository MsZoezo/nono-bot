package cli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/bot"
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

		bot, err := bot.New(token)

		if err != nil {
			log.Fatal("Bot couldn't initalize..")
		}

		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc

		bot.Close()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
