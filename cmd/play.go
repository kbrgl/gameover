package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	playCmd = &cobra.Command{
		Use:   "play [game]",
		Short: "Play a game",
		Args:  cobra.ExactArgs(1),
		Run:   play,
	}
)

func init() {
	rootCmd.AddCommand(playCmd)
}

func play(cmd *cobra.Command, args []string) {
	games, err := FetchGames()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = PlayGame(FindGame(args[0], games))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
