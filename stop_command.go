package main

import (
	"fmt"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type stopCommand struct{}

func (command stopCommand) Run() error {
	if err := systemutils.KillBackgroundProcess(appName); err != nil {
		return fmt.Errorf("unable to kill the background process: %w", err)
	}

	return nil
}
