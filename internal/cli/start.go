package cli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MsZoezo/nono-bot/internal/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [token]",
	Short: "Start nono bot",
	Run: func(_ *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing token argument, aborting.")
			return
		}

		dg, err := discordgo.New("Bot " + args[0])

		if err != nil {
			fmt.Println("Error creating discord session, aborting.")
			return
		}

		dg.Identify.Intents = discordgo.IntentGuildMessages

		dg.AddHandler(bot.MessageCreateHandler)

		err = dg.Open()

		if err != nil {
			fmt.Println("Error opening connection, aborting.")
			return
		}

		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc

		dg.Close()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
