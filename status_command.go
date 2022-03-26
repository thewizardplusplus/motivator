package main

import (
	"fmt"

	ps "github.com/mitchellh/go-ps"
)

type statusCommand struct{}

func (command statusCommand) Run() error {
	processes, err := ps.Processes()
	if err != nil {
		return fmt.Errorf("unable to get a process list: %w", err)
	}

	var instanceCount int
	for _, process := range processes {
		if process.Executable() == "motivator" {
			instanceCount++
		}
	}

	if instanceCount > 1 {
		fmt.Println("motivator is running in background")
	} else {
		fmt.Println("motivator is not running in background")
	}

	return nil
}
