package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	leaderboardCmd = &cobra.Command{
		Use:   "friend",
		Short: "Friend someone by their username",
		Run:   friend,
	}
)

func init() {
	rootCmd.AddCommand(leaderboardCmd)
}

func friend(cmd *cobra.Command, args []string) {
	fmt.Println("friending isn't implemented yet")
}
