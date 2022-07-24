package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize ~/.gameover.",
		Run:   initialize,
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

func initialize(cmd *cobra.Command, args []string) {
	tmp, err := os.MkdirTemp("", "gameover")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.RemoveAll(gameoverDir(""))
	_, err = git.PlainClone(tmp, false, &git.CloneOptions{
		URL: "https://github.com/kbrgl/gameover.git",
	})
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	err = os.Rename(filepath.Join(tmp, ".gameover"), gameoverDir(""))
	if err != nil {
		fmt.Printf("failed to rename: %s\n", err)
		os.Exit(1)
	}
	err = os.RemoveAll(tmp)
	if err != nil {
		fmt.Printf("failed to delete directory: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Initialized gameover in", gameoverDir(""))
	fmt.Println("Add 'source ~/.gameover/env' to your ~/.profile or whatever shell init script you use")
}
