package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ghonijee/php-version-manager/utils"
	"github.com/urfave/cli/v2"
)

func List(ctx *cli.Context) error {
	pathInstalled := utils.PathInstalled()
	entries, err := os.ReadDir(pathInstalled)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("PHP Installed:")
	for _, dir := range entries {
		if strings.Contains(dir.Name(), "php") && dir.IsDir() {
			commandList := fmt.Sprintf("%s%s", pathInstalled, dir.Name())
			dirPhp, err := exec.Command("ls", commandList).Output()
			if err != nil {
				log.Panicln(err)
			}
			fmt.Printf("php %s \n", string(dirPhp)[:3])
		}
	}

	return err
}
