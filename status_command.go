package main

import (
	"fmt"
)

type statusCommand struct{}

func (command statusCommand) Run() error {
	backgroundProcess, err := findBackgroundProcess(appName)
	if err != nil {
		return fmt.Errorf("unable to find a background process: %w", err)
	}

	if backgroundProcess != nil {
		fmt.Printf("%s is running in background\n", appName)
	} else {
		fmt.Printf("%s is not running in background\n", appName)
	}

	return nil
}
