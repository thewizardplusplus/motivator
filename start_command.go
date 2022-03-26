package main

import (
	"fmt"
	"os"
	"os/exec"
)

type startCommand struct {
	configurableCommand
}

func (command startCommand) Run() error {
	if err := killBackgroundProcess(appName); err != nil {
		return fmt.Errorf("unable to kill a previous background process: %w", err)
	}

	// run a second instance of itself in background
	backgroundInstance := exec.Command(
		os.Args[0],
		"foreground",
		"--config",
		command.ConfigPath,
	)
	if err := backgroundInstance.Start(); err != nil {
		return fmt.Errorf(
			"unable to run a second instance of itself in background: %w",
			err,
		)
	}

	return nil
}
