package cmd

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

var (
	// GameoverDir is the directory where games are installed.
	GameoverDir = "~/.gameover"
	// BinariesDir is the directory where binaries are installed.
	BinariesDir = "~/.gameover/bin"
)

func init() {
	GameoverDir, _ = homedir.Expand(GameoverDir)
	BinariesDir, _ = homedir.Expand(BinariesDir)
	os.MkdirAll(BinariesDir, os.ModePerm)
}

func gameoverDir(path string) string {
	return filepath.Join(GameoverDir, path)
}

func binariesDir(path string) string {
	return filepath.Join(BinariesDir, path)
}
