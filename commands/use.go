package commands

import (
	"errors"

	"github.com/urfave/cli/v2"
)

func Use(ctx *cli.Context) error {
	newVersion := ctx.Args().First()

	if len(newVersion) == 0 {
		return errors.New("PHP version use not parse")
	}
	// // get Path PHP formula on homebrew
	// brewPath := utils.BrewPhpInstalled()
	// // Current PHP active path bin
	// getPhpBinCommand := fmt.Sprintf("ls -d %sphp*/%s*/bin", brewPath, currentVersion)
	// currentVersionBinPath := utils.Cli(getPhpBinCommand)

	// // Target PHP new path bin
	// targetPhpBinCommand := fmt.Sprintf("ls -d %sphp*/%s*/bin", brewPath, newVersion)
	// targetVersionBinPath := utils.Cli(targetPhpBinCommand)

	// exportCommand := fmt.Sprintf("export PATH=%s:$PATH", targetVersionBinPath.Removeln())
	// utils.Cli(exportCommand)
	// exec.Command("sh", "c", exportCommand)
	// fmt.Printf("PHP use: %s \n\n", newVersion)
	// fmt.Println(exportCommand)
	return nil
}
