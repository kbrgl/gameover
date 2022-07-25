package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	leaderboardCmd = &cobra.Command{
		Use:   "friend",
		Short: "Friend someone by their username",
		Args:  cobra.ExactArgs(1),
		Run:   friend,
	}
)

func init() {
	rootCmd.AddCommand(leaderboardCmd)
}

func friend(cmd *cobra.Command, args []string) {
	fr := args[0]

	form := url.Values{}
	form.Add("friend", fr)

	req, err := http.NewRequest("POST", os.Getenv("GAMEOVER_SERVER_URL")+"/friend", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("could not friend:", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("could not friend:", resp.Status)
		os.Exit(1)
	}

	fmt.Println("Friended")
}
