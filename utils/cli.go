package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CliResult struct {
	Value string
	Error string
}

// Remove new line from string Value
func (val CliResult) Removeln() string {
	return strings.ReplaceAll(val.Value, "\n", "")
}

func Cli(arg string) CliResult {
	output, err := exec.Command("bash", "-c", arg).Output()
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		os.Exit(1)
	}

	result := CliResult{
		Value: string(output),
	}
	return result
}
func CliNew(name string, arg ...string) CliResult {
	output, err := exec.Command(name, arg...).Output()
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		os.Exit(1)
	}

	result := CliResult{
		Value: string(output),
	}
	return result
}
