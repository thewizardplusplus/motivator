package main

import (
	"fmt"
	"os"
	"os/exec"
)

type startCommand struct {
	ConfigPath string `kong:"required,short='c',name='config',default='config.json',help='Config path.'"` // nolint: lll
}

func (command startCommand) Run() error {
	if err := killBackgroundProcess("motivator"); err != nil {
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
