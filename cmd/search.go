package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	searchCmd = &cobra.Command{
		Use:   "search [game]",
		Short: "Search for a game",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			games, err := FetchGames()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			query := ""
			if len(args) > 0 {
				query = args[0]
			}
			results := FilterGames(query, games)
			resStyle := (lipgloss.NewStyle().
				Foreground(lipgloss.Color("15")).
				Background(lipgloss.Color("12")).
				PaddingLeft(1).PaddingRight(1).
				Bold(true))
			fmt.Print(resStyle.Render("Results"))
			faint := lipgloss.NewStyle().Faint(true)
			fmt.Printf(faint.Render(" searching %d game(s):")+"\n", len(results))
			for _, g := range results {
				fmt.Printf("  %s %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render(g.Name), g.Repo)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(searchCmd)
}
