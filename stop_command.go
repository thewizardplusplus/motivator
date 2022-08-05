package main

import (
	"fmt"
	"os"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type stopCommand struct{}

func (command stopCommand) Run() error {
	_, executableName, err := systemutils.ExecutableOfForegroundProcess()
	if err != nil {
		const message = "unable to get the name of the executable " +
			"of the foreground process: %w"
		return fmt.Errorf(message, err)
	}

	if err := systemutils.KillBackgroundProcess(
		executableName,
		os.Getpid(),
	); err != nil {
		return fmt.Errorf("unable to kill the background process: %w", err)
	}

	return nil
}
