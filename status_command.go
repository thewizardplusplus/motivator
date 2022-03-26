package main

import (
	"fmt"
	"os"

	ps "github.com/mitchellh/go-ps"
)

type statusCommand struct{}

func (command statusCommand) Run() error {
	processes, err := ps.Processes()
	if err != nil {
		return fmt.Errorf("unable to get a process list: %w", err)
	}

	var isRunning bool
	currentPID := os.Getpid()
	for _, process := range processes {
		if process.Executable() == "motivator" && process.Pid() != currentPID {
			isRunning = true
			break
		}
	}

	if isRunning {
		fmt.Println("motivator is running in background")
	} else {
		fmt.Println("motivator is not running in background")
	}

	return nil
}
