package main

import (
	"fmt"
	"os"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type stopCommand struct {
	executableInfoCommand `kong:"-"`
}

func (command stopCommand) Run() error {
	if err := systemutils.KillBackgroundProcess(
		command.ExecutableInfo.Name,
		os.Getpid(),
	); err != nil {
		return fmt.Errorf("unable to kill the background process: %w", err)
	}

	return nil
}
