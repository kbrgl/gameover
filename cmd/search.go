package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	searchCmd = &cobra.Command{
		Use:   "search [game]",
		Short: "Search for a game",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			games, err := FetchGames()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			results := FilterGames(args[0], games)
			for _, g := range results {
				fmt.Printf("%s\n", g)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(searchCmd)
}
