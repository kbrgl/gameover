package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	playCmd = &cobra.Command{
		Use:   "play [game]",
		Short: "Play a game",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			games, err := FetchGames()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = Play(FindGame(args[0], games))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(playCmd)
}

func Play(g *Game) error {
	bin := binariesDir(g.Name)
	cmd := exec.Command(bin)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
