package main

import (
	"fmt"
	"os"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type statusCommand struct {
	executableInfoCommand `kong:"-"`
}

func (command statusCommand) Run() error {
	backgroundProcess, err :=
		systemutils.FindBackgroundProcess(command.ExecutableInfo.Name, os.Getpid())
	if err != nil {
		return fmt.Errorf("unable to find the background process: %w", err)
	}

	if backgroundProcess != nil {
		fmt.Printf("%s is running in the background\n", command.ExecutableInfo.Name)
	} else {
		fmt.Printf(
			"%s is not running in the background\n",
			command.ExecutableInfo.Name,
		)
	}

	return nil
}
