package cmd

import (
	"bytes"
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
	registerCmd = &cobra.Command{
		Use:   "register",
		Short: "Register a new user",
		Run:   register,
	}
)

func init() {
	rootCmd.AddCommand(registerCmd)
}

func register(cmd *cobra.Command, args []string) {
	var username string
	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	password, _ := term.ReadPassword(syscall.Stdin)
	fmt.Println()

	fmt.Print("Confirm password: ")
	confirm, _ := term.ReadPassword(syscall.Stdin)
	fmt.Println()
	if !bytes.Equal(password, confirm) {
		fmt.Println("Passwords do not match")
		os.Exit(1)
	}

	var email string
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	fmt.Println()

	form := url.Values{}
	form.Add("username", username)
	form.Add("password", string(password))
	form.Add("email", email)

	resp, err := http.Post(
		os.Getenv("GAMEOVER_SERVER_URL")+"/register",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("could not register:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var result map[string]string
		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Println("could not register:", result["message"])
		os.Exit(1)
	}
	fmt.Println("user created, please log in.")
}
