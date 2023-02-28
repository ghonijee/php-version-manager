package utils

import (
	"log"
	"os/exec"
	"strings"
)

func BrewPhpInstalled() string {
	brewPrefix, err := exec.Command("brew", "--prefix").Output()
	if err != nil {
		log.Panicln(err)
	}
	path := []string{string(brewPrefix), "Cellar/"}
	fullPath := strings.Join(path, "/")
	return strings.ReplaceAll(fullPath, "\n", "")
}

func PathBin() string {
	return ""
}
