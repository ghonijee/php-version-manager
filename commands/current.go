package commands

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/urfave/cli/v2"
)

func Current(ctx *cli.Context) error {
	result, err := exec.Command("php", "-v").Output()
	if err != nil {
		log.Panicln(err)
		return err
	}
	version := findVersionText(string(result))
	fmt.Printf("PHP version active: %s ", version)
	return nil
}

func findVersionText(text string) string {
	regex, err := regexp.Compile("[0-9]+.[0-9]+")
	if err != nil {
		log.Panicln(err)
		return ""
	}
	version := regex.FindString(text)
	return version
}
