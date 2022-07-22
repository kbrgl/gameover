package cmd

import (
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
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
		games []*Game
	}
	_, err = toml.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("could not decode games")
	}
	return config.games, nil
}
