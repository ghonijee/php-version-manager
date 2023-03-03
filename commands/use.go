package commands

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/ghonijee/php-version-manager/utils"
	"github.com/urfave/cli/v2"
)

func Use(ctx *cli.Context) error {
	newVersion := ctx.Args().First()

	if len(newVersion) == 0 {
		return errors.New("PHP version use not parse")
	}

	// Open file bash config for pvm
	var file, err = os.OpenFile(getFilePath(), os.O_RDWR, 0644)
	utils.ErrCheck(err)
	defer file.Close()
	// Read content of file
	content, err := os.ReadFile(getFilePath())
	utils.ErrCheck(err)

	// get Path PHP formula on homebrew
	brewPath := utils.BrewPhpInstalled()
	resultPhpVersion := utils.CliNew("php", "-v")
	currentVersion := findVersionText(resultPhpVersion.Value)
	// Current PHP active path bin
	getPhpBinCommand := fmt.Sprintf("ls -d %sphp*/%s*/bin", brewPath, currentVersion)
	currentVersionBinPath := utils.Cli(getPhpBinCommand)

	// Target PHP new path bin
	targetPhpBinCommand := fmt.Sprintf("ls -d %sphp*/%s*/bin", brewPath, newVersion)
	targetVersionBinPath := utils.Cli(targetPhpBinCommand)

	if currentVersion == newVersion {
		fmt.Printf("PHP version %s already active", newVersion)
		return nil
	}

	var checkPath = bytes.Contains(content, []byte(currentVersionBinPath.Removeln()))
	if !checkPath {
		defaultConfig := fmt.Sprintf("\n\nexport PATH=%s:$PATH", currentVersionBinPath.Removeln())
		content = append(content, defaultConfig...)
	}
	content = bytes.Replace(content, []byte(currentVersionBinPath.Removeln()), []byte(targetVersionBinPath.Removeln()), -1)
	_, err = file.WriteAt(content, 0)
	utils.ErrCheck(err)
	err = file.Sync()
	utils.ErrCheck(err)

	// Source bash for apply change version
	targetCommand := fmt.Sprintf("export PATH=%s:$PATH", targetVersionBinPath.Removeln())
	fmt.Println(targetCommand)
	utils.Cli("source " + getFilePath())
	utils.CliNew("bash", "-c", targetCommand)
	utils.CliNew("bash", "source", getFilePath())

	fmt.Println("Use PHP version: " + newVersion)
	return nil
}

func getFilePath() string {
	result := utils.Cli("echo $HOME")
	bashFileLocation := fmt.Sprintf("%s/%s", result.Removeln(), fileName)
	return bashFileLocation
}
