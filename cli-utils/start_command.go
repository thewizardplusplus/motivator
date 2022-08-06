package cliutils

import (
	"fmt"
	"io"

	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type StartCommand struct {
	StopCommand
}

func (command StartCommand) Run(
	backgroundProcessArguments []string,
	stderrWriter io.Writer,
	stderrEndMark string,
) error {
	if err := command.StopCommand.Run(); err != nil {
		return fmt.Errorf("unable to run the `stop` command: %w", err)
	}

	if err := systemutils.StartBackgroundProcess(
		command.ExecutableInfo.Path,
		backgroundProcessArguments,
		stderrWriter,
		stderrEndMark,
	); err != nil {
		return fmt.Errorf("unable to start the background process: %w", err)
	}

	return nil
}
