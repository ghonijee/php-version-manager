package main

import (
	"log"
	"os"

	"github.com/ghonijee/php-version-manager/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "PHP Version Manager",
		Usage:     "A cli tool for manage PHP versions",
		UsageText: "pvm command [command options] [arguments...]",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "Show list PHP installed",
				Aliases: []string{"ls"},
				Action:  commands.List,
			},
			{
				Name:    "current",
				Usage:   "Show PHP version active",
				Aliases: []string{"c"},
				Action:  commands.Current,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
