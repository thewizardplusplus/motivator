package cliutils

import (
	"fmt"
	"os"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type StopCommand struct {
	ExecutableInfoCommand `kong:"-"`
}

func (command StopCommand) Run() error {
	if err := systemutils.KillBackgroundProcess(
		command.ExecutableInfo.Name,
		os.Getpid(),
	); err != nil {
		return fmt.Errorf("unable to kill the background process: %w", err)
	}

	return nil
}
