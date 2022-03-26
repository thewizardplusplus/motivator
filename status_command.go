package main

import (
	"fmt"
)

type statusCommand struct{}

func (command statusCommand) Run() error {
	backgroundProcess, err := findBackgroundProcess("motivator")
	if err != nil {
		return fmt.Errorf("unable to find a background process: %w", err)
	}

	if backgroundProcess != nil {
		fmt.Println("motivator is running in background")
	} else {
		fmt.Println("motivator is not running in background")
	}

	return nil
}
