package cmd

import "path/filepath"

var (
	// GameoverDir is the directory where games are installed.
	GameoverDir = "~/.gameover"
	// BinariesDir is the directory where binaries are installed.
	BinariesDir = "~/.gameover/bin"
)

func gameoverDir(path string) string {
	return filepath.Join(GameoverDir, path)
}

func binariesDir(path string) string {
	return filepath.Join(BinariesDir, path)
}
