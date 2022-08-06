package cliutils

import (
	"fmt"
	"os"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type StatusCommand struct {
	ExecutableInfoCommand `kong:"-"`
}

func (command StatusCommand) Run() error {
	backgroundProcess, err :=
		systemutils.FindBackgroundProcess(command.ExecutableInfo.Name, os.Getpid())
	if err != nil {
		return fmt.Errorf("unable to find the background process: %w", err)
	}

	var status string
	if backgroundProcess != nil {
		status = "is running"
	} else {
		status = "is not running"
	}

	fmt.Printf("%s status: %s\n", command.ExecutableInfo.Name, status)
	return nil
}
