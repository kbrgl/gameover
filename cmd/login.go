package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login to the application",
		Run:   login,
	}
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

func saveToken(token string) error {
	f, err := os.Create(gameoverDir("token"))
	if err != nil {
		return err
	}

	_, err = f.WriteString(token)
	return err
}

func login(cmd *cobra.Command, args []string) {
	var username string
	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	password, _ := term.ReadPassword(syscall.Stdin)

	form := url.Values{}
	form.Add("username", username)
	form.Add("password", string(password))

	r, err := http.Post(
		os.Getenv("GAMEOVER_SERVER_URL")+"/login",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("could not login:", err)
		os.Exit(1)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		fmt.Println("could not login:", r.Status)
		os.Exit(1)
	}

	var result map[string]string
	err = json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		fmt.Println("could not login:", err)
		os.Exit(1)
	}

	fmt.Println("Saving token:", result["token"])
	err = saveToken(result["token"])
	if err != nil {
		fmt.Println("could not save token:", err)
		os.Exit(1)
	}
}
