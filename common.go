package main

import (
	"fmt"
	"os"

	ps "github.com/mitchellh/go-ps"
)

func findBackgroundProcess(executableName string) (ps.Process, error) {
	processes, err := ps.Processes()
	if err != nil {
		return nil, fmt.Errorf("unable to get a process list: %w", err)
	}

	var backgroundProcess ps.Process
	currentPID := os.Getpid()
	for _, process := range processes {
		if process.Executable() == executableName && process.Pid() != currentPID {
			backgroundProcess = process
			break
		}
	}

	return backgroundProcess, nil
}

func killBackgroundProcess(executableName string) error {
	backgroundProcess, err := findBackgroundProcess(executableName)
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
