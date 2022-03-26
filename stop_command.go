package main

import (
	"fmt"
	"os"
)

type stopCommand struct{}

func (command stopCommand) Run() error {
	backgroundProcess, err := findBackgroundProcess()
	if err != nil {
		return fmt.Errorf("unable to find a background process: %w", err)
	}
	if backgroundProcess == nil {
		return nil
	}

	osBackgroundProcess, err := os.FindProcess(backgroundProcess.Pid())
	if err != nil {
		return fmt.Errorf("unable to find a background process by its PID: %w", err)
	}

	if err := osBackgroundProcess.Kill(); err != nil {
		return fmt.Errorf("unable to kill a background process: %w", err)
	}

	return nil
}
