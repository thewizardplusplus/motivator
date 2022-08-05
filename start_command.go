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
	executablePath, executableName, err :=
		systemutils.ExecutableOfForegroundProcess()
	if err != nil {
		const message = "unable to get the path and the name of the executable " +
			"of the foreground process: %w"
		return fmt.Errorf(message, err)
	}

	if err := systemutils.KillBackgroundProcess(
		executableName,
		os.Getpid(),
	); err != nil {
		return fmt.Errorf("unable to kill the background process: %w", err)
	}

	if err := systemutils.StartBackgroundProcess(
		executablePath,
		[]string{"foreground", "--config", command.ConfigPath},
		os.Stderr,
		markOfShowingStart,
	); err != nil {
		return fmt.Errorf("unable to start the background process: %w", err)
	}

	return nil
}
