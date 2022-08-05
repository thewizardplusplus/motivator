package main

import (
	"fmt"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

const (
	markOfShowingStart = "start showing notifications"
)

type configurableCommand struct {
	ConfigPath string `kong:"required,short='c',name='config',default='config.json',help='Config path.'"` // nolint: lll
}

type executableInfoCommand struct {
	ExecutableInfo systemutils.ExecutableInfo
}

func (command *executableInfoCommand) AfterApply() error {
	executableInfo, err := systemutils.ExecutableOfForegroundProcess()
	if err != nil {
		const message = "unable to get the name of the executable " +
			"of the foreground process: %w"
		return fmt.Errorf(message, err)
	}

	command.ExecutableInfo = executableInfo
	return nil
}
