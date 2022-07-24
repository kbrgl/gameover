package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/go-git/go-git/v5"
	"github.com/kbrgl/fuzzy"
)

type Game struct {
	Name string `toml:"name"`
	Repo string `toml:"repo"`
}

func (g *Game) String() string {
	return fmt.Sprintf("%s (%s)", g.Name, g.Repo)
}

func FindGame(name string, games []*Game) *Game {
	for _, g := range games {
		if g.Name == name {
			return g
		}
	}
	return nil
}

func FilterGames(query string, games []*Game) []*Game {
	result := make([]*Game, 0)
	for _, g := range games {
		if fuzzy.MatchFold(g.Name, query) {
			result = append(result, g)
		}
	}
	return result
}

func FetchGames() ([]*Game, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/kbrgl/gameover/master/games.toml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("could not fetch games: %s", resp.Status)
	}
	var config struct {
		Games []*Game `toml:"games"`
	}
	_, err = toml.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("could not decode games")
	}
	return config.Games, nil
}

func buildAndAdd(fp string) error {
	os.Chdir(fp)
	cmdDownload := exec.Command("go", "mod", "download")
	cmdDownload.Stdout = os.Stdout
	cmdDownload.Stderr = os.Stderr
	err := cmdDownload.Run()
	if err != nil {
		return err
	}
	cmdBuild := exec.Command("go", "build", "-o", filepath.Join("../bin", filepath.Base(fp)))
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	return cmdBuild.Run()
}

func InstallGame(g *Game) error {
	fp := gameoverDir(g.Name)
	os.RemoveAll(fp)

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

func PlayGame(g *Game) error {
	bin := binariesDir(g.Name)
	cmd := exec.Command(bin)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
