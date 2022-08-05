package main

import (
	"fmt"
	"os"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type statusCommand struct{}

func (command statusCommand) Run() error {
	_, executableName, err := systemutils.ExecutableOfForegroundProcess()
	if err != nil {
		const message = "unable to get the name of the executable " +
			"of the foreground process: %w"
		return fmt.Errorf(message, err)
	}

	backgroundProcess, err :=
		systemutils.FindBackgroundProcess(executableName, os.Getpid())
	if err != nil {
		return fmt.Errorf("unable to find the background process: %w", err)
	}

	if backgroundProcess != nil {
		fmt.Printf("%s is running in the background\n", executableName)
	} else {
		fmt.Printf("%s is not running in the background\n", executableName)
	}

	return nil
}
