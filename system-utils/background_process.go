package systemutils

import (
	"fmt"
	"os"

	ps "github.com/mitchellh/go-ps"
)

func FindBackgroundProcess(executableName string) (ps.Process, error) {
	processes, err := ps.Processes()
	if err != nil {
		return nil, fmt.Errorf("unable to get the process list: %w", err)
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

func KillBackgroundProcess(executableName string) error {
	backgroundProcess, err := FindBackgroundProcess(executableName)
	if err != nil {
		return fmt.Errorf("unable to find the background process by name: %w", err)
	}
	if backgroundProcess == nil {
		return nil
	}

	osBackgroundProcess, err := os.FindProcess(backgroundProcess.Pid())
	if err != nil {
		return fmt.Errorf("unable to find the background process by PID: %w", err)
	}

	if err := osBackgroundProcess.Kill(); err != nil {
		return fmt.Errorf("unable to kill the background process: %w", err)
	}

	return nil
}
