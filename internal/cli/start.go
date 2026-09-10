package cli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"charm.land/log/v2"
	"github.com/MsZoezo/nono-bot/internal/bot"
	"github.com/bwmarrin/discordgo"
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

		dg, err := discordgo.New(token)

		if err != nil {
			fmt.Println("Error creating discord session, aborting.")
			return
		}

		log.Debug("Succesfully created discord session.")

		dg.Identify.Intents = discordgo.IntentGuildMessages

		dg.AddHandler(bot.MessageCreateHandler)
		dg.AddHandler(bot.ConnectHandler)
		dg.AddHandler(bot.DisconnectHandler)

		err = dg.Open()

		if err != nil {
			fmt.Println("Error opening connection, aborting.")
			return
		}

		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc

		log.Info("Shutting down..")

		dg.Close()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
