package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var (
	installCmd = &cobra.Command{
		Use:   "install [game]",
		Short: "Install a game",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			games, err := FetchGames()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = Install(FindGame(args[0], games))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
}

func buildAndAdd(fp string) error {
	cmd := exec.Command("go", "build", "-o", binariesDir(filepath.Base(fp)), fp)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Install(g *Game) error {
	fp := gameoverDir(g.Name)
	_, err := git.PlainClone(fp, false, &git.CloneOptions{
		URL: g.Repo,
	})
	if err != nil {
		if err == git.ErrRepositoryAlreadyExists {
			return fmt.Errorf("game already installed: %s", g.Name)
		}
		return err
	}

	err = buildAndAdd(fp)
	if err != nil {
		return err
	}

	return nil
}
