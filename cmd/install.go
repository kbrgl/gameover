package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	installCmd = &cobra.Command{
		Use:   "install [game]",
		Short: "Install a game",
		Args:  cobra.ExactArgs(1),
		Run:   install,
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
}

func install(cmd *cobra.Command, args []string) {
	games, err := FetchGames()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = InstallGame(FindGame(args[0], games))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
