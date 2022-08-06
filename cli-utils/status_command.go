package cliutils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

var (
	blueColor  = color.New(color.FgBlue, color.Bold).SprintFunc()
	greenColor = color.New(color.FgGreen, color.Bold).SprintFunc()
	redColor   = color.New(color.FgRed, color.Bold).SprintFunc()
)

type StatusCommand struct {
	ExecutableInfoCommand `kong:"-"`
}

// Help implements the `kong.HelpProvider` interface.
func (command StatusCommand) Help() string {
	return "This command supports the NO_COLOR environment variable " +
		"that disables colorful output."
}

func (command StatusCommand) Run() error {
	backgroundProcess, err :=
		systemutils.FindBackgroundProcess(command.ExecutableInfo.Name, os.Getpid())
	if err != nil {
		return fmt.Errorf("unable to find the background process: %w", err)
	}

	var status string
	if backgroundProcess != nil {
		status = greenColor("is running")
	} else {
		status = redColor("is not running")
	}

	fmt.Printf("%s status: %s\n", blueColor(command.ExecutableInfo.Name), status)
	return nil
}
