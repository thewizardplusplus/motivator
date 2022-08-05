package main

import (
	"fmt"
	"os"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type startCommand struct {
	stopCommand
	configurableCommand
}

func (command startCommand) Run() error {
	if err := command.stopCommand.Run(); err != nil {
		return fmt.Errorf("unable to run the `stop` command: %w", err)
	}

	if err := systemutils.StartBackgroundProcess(
		command.ExecutableInfo.Path,
		[]string{"foreground", "--config", command.ConfigPath},
		os.Stderr,
		markOfShowingStart,
	); err != nil {
		return fmt.Errorf("unable to start the background process: %w", err)
	}

	return nil
}
