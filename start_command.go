package main

import (
	"fmt"
	"os"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type startCommand struct {
	configurableCommand
}

func (command startCommand) Run() error {
	executableInfo, err := systemutils.ExecutableOfForegroundProcess()
	if err != nil {
		const message = "unable to get the path and the name of the executable " +
			"of the foreground process: %w"
		return fmt.Errorf(message, err)
	}

	if err := systemutils.KillBackgroundProcess(
		executableInfo.Name,
		os.Getpid(),
	); err != nil {
		return fmt.Errorf("unable to kill the background process: %w", err)
	}

	if err := systemutils.StartBackgroundProcess(
		executableInfo.Path,
		[]string{"foreground", "--config", command.ConfigPath},
		os.Stderr,
		markOfShowingStart,
	); err != nil {
		return fmt.Errorf("unable to start the background process: %w", err)
	}

	return nil
}
