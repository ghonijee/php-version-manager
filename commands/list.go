package commands

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ghonijee/php-version-manager/utils"
	"github.com/urfave/cli/v2"
)

func List(ctx *cli.Context) error {
	BrewPhpInstalled := utils.BrewPhpInstalled()
	entries, err := os.ReadDir(BrewPhpInstalled)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("PHP Installed: ")
	for _, dir := range entries {
		if strings.Contains(dir.Name(), "php") && dir.IsDir() {
			commandList := fmt.Sprintf("ls %s%s", BrewPhpInstalled, dir.Name())
			dirPhp := utils.Cli(commandList)
			fmt.Printf("php %s \n", dirPhp.Removeln()[:3])
		}
	}

	return err
}
