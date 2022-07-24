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
	if err := systemutils.KillBackgroundProcess(appName); err != nil {
		return fmt.Errorf("unable to kill the background process: %w", err)
	}

	if err := systemutils.StartBackgroundProcess(
		os.Args[0],
		[]string{"foreground", "--config", command.ConfigPath},
		os.Stderr,
		markOfShowingStart,
	); err != nil {
		return fmt.Errorf("unable to start the background process: %w", err)
	}

	return nil
}
