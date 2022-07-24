package main

import (
	"fmt"
	"os"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

const (
	appName            = "motivator"
	markOfShowingStart = "start showing notifications"
)

type configurableCommand struct {
	ConfigPath string `kong:"required,short='c',name='config',default='config.json',help='Config path.'"` // nolint: lll
}

func killBackgroundProcess(executableName string) error {
	backgroundProcess, err := systemutils.FindBackgroundProcess(executableName)
	if err != nil {
		return fmt.Errorf("unable to find the background process: %w", err)
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
