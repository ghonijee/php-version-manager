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
				Name:    "use",
				Usage:   "Change PHP version active",
				Aliases: []string{"u"},
				Action:  commands.Use,
			},
			{
				Name:    "current",
				Usage:   "Show PHP version active",
				Aliases: []string{"c"},
				Action:  commands.Current,
			},
			{
				Name:    "init",
				Usage:   "Initial setup for pvm",
				Aliases: []string{"in"},
				Action:  commands.Init,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
