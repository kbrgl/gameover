package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	binCmd = &cobra.Command{
		Use:   "bin",
		Short: "Print the path to the binaries directory",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(binariesDir(""))
		},
	}
)

func init() {
	rootCmd.AddCommand(binCmd)
}
