package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type startCommand struct {
	configurableCommand
}

func (command startCommand) Run() error {
	if err := killBackgroundProcess(appName); err != nil {
		return fmt.Errorf("unable to kill a previous background process: %w", err)
	}

	// run a second instance of itself in background
	backgroundInstance := exec.Command(
		os.Args[0],
		"foreground",
		"--config",
		command.ConfigPath,
	)
	stderr, err := backgroundInstance.StderrPipe()
	if err != nil {
		return fmt.Errorf("unable to get a stderr pipe: %w", err)
	}
	defer stderr.Close()

	if err := backgroundInstance.Start(); err != nil {
		return fmt.Errorf(
			"unable to run a second instance of itself in background: %w",
			err,
		)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, markOfShowingStart) {
			break
		}

		fmt.Fprintln(os.Stderr, line)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("unable to read a stderr of a background instance: %w", err)
	}

	return nil
}
