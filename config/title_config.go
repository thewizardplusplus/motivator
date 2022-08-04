package config

import (
	"strings"

	"github.com/thewizardplusplus/motivator/entities"
)

type TitleConfig struct {
	HideAppName         bool
	UseOriginalTaskName bool
}

func (config TitleConfig) Title(appName string, task entities.Task) string {
	taskName := task.SelectedName()
	if config.UseOriginalTaskName {
		taskName = task.OriginalName
	}

	var titleParts []string
	if !config.HideAppName {
		titleParts = append(titleParts, appName)
	}
	if taskName != "" {
		titleParts = append(titleParts, taskName)
	}

	return strings.Join(titleParts, " | ")
}
