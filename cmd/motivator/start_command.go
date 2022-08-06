package main

import (
	"os"

	cliutils "github.com/thewizardplusplus/motivator/cli-utils"
)

type startCommandWrapper struct {
	cliutils.StartCommand
	configurableCommand
}

func (command startCommandWrapper) Run() error {
	return command.StartCommand.Run(
		[]string{"foreground", "--config", command.ConfigPath},
		os.Stderr,
		notificationDisplayStartMark,
	)
}
