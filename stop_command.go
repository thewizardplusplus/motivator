package main

import (
	"fmt"
)

type stopCommand struct{}

func (command stopCommand) Run() error {
	if err := killBackgroundProcess(appName); err != nil {
		return fmt.Errorf("unable to kill a background process: %w", err)
	}

	return nil
}
