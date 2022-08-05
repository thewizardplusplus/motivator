package cliutils

import (
	"fmt"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type ExecutableInfoCommand struct {
	ExecutableInfo systemutils.ExecutableInfo
}

// AfterApply implements the `kong.AfterApply` interface.
func (command *ExecutableInfoCommand) AfterApply() error {
	executableInfo, err := systemutils.ExecutableOfForegroundProcess()
	if err != nil {
		const message = "unable to get the name of the executable " +
			"of the foreground process: %w"
		return fmt.Errorf(message, err)
	}

	command.ExecutableInfo = executableInfo
	return nil
}
