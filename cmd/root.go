package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gameover",
	Short: `   ___                                         
  / _ \__ _ _ __ ___   ___  _____   _____ _ __ 
 / /_\/ _' | '_ ' _ \ / _ \/ _ \ \ / / _ \ '__|
/ /_\\ (_| | | | | | |  __/ (_) \ V /  __/ |   
\____/\__,_|_| |_| |_|\___|\___/ \_/ \___|_|

Play text-based games right in your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var token string

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

func Execute() {
	dat, err := os.ReadFile(gameoverDir("token"))
	if err != nil {
		fmt.Println("could not read login token:", err)
		os.Exit(1)
	}
	token = string(dat)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
