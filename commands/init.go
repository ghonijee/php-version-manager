package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/ghonijee/php-version-manager/utils"
	"github.com/urfave/cli/v2"
)

var phpTap = "shivammathur/php"
var BREW_DISABLE_AUTO_CLEANUP = "HOMEBREW_NO_INSTALL_CLEANUP=1"
var filePath = "/Users/yunus-floo/.pvm_zsh.zsh"

func Init(ctx *cli.Context) error {
	var result utils.CliResult
	// Step by Step
	// 1. whether brew is installed
	result = utils.Cli("which brew")
	if !strings.Contains(result.Value, "homebrew") {
		return cli.Exit("Homebrew is not installed, Visit https://brew.sh/ for instructions on installing Homebrew.", 1)
	}
	// 2. whether brew have tap from  'homebrew/homebrew-core' or 'shivammathur/php'
	result = utils.Cli("brew tap")
	if !strings.Contains(result.Value, phpTap) {
		fmt.Printf("PHP is not tapped on Homebrew, tapped now via brew tap \"%s\" \n", phpTap)
		utils.Cli(BREW_DISABLE_AUTO_CLEANUP + " brew tap " + phpTap)
	}

	// 3. whether php installed
	result = utils.CliNew("brew", "list", "--formula")
	if !strings.Contains(result.Value, "php") {
		fmt.Printf("PHP is not installed, installing now via brew tap \"%s\" \n", phpTap)
		utils.Cli(BREW_DISABLE_AUTO_CLEANUP + " brew install php")
	}
	file, err := os.Create(filePath)
	utils.ErrCheck(err)
	defer file.Close()

	// value := []byte("# Generate by PHP Version Manager (pvm) \n\n")
	_, err = file.WriteString("# Generate by PHP Version Manager (pvm) \n\n")
	utils.ErrCheck(err)

	return nil
}
