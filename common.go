package main

import (
	"fmt"
	"os"

	ps "github.com/mitchellh/go-ps"
)

func findBackgroundProcess() (ps.Process, error) {
	processes, err := ps.Processes()
	if err != nil {
		return nil, fmt.Errorf("unable to get a process list: %w", err)
	}

	var backgroundProcess ps.Process
	currentPID := os.Getpid()
	for _, process := range processes {
		if process.Executable() == "motivator" && process.Pid() != currentPID {
			backgroundProcess = process
			break
		}
	}

	return backgroundProcess, nil
}
