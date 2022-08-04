package config

import (
	"path/filepath"

	"github.com/thewizardplusplus/motivator/entities"
)

type Config struct {
	TitleConfig

	Icon      string
	Tasks     []entities.Task
	Variables map[string]string
}

func (config Config) PrepareTasks(
	generatedNamePrefix string,
	basicIconPath string,
) []entities.Task {
	var tasks []entities.Task
	taskNameGenerator := entities.NewNameGenerator(generatedNamePrefix)
	for taskIndex, task := range config.Tasks {
		if len(task.Phrases) == 0 {
			continue
		}

		task.OriginalName = task.Name
		task.Name = taskNameGenerator.GenerateName(taskIndex, task.Name)

		for phraseIndex, phrase := range task.Phrases {
			iconPath := entities.CoalesceStrings(phrase.Icon, task.Icon, config.Icon)
			if iconPath != "" && !filepath.IsAbs(iconPath) {
				iconPath = filepath.Join(basicIconPath, iconPath)
			}

			task.Phrases[phraseIndex] = entities.Phrase{
				Icon: iconPath,
				Text: phrase.ExpandText(config.Variables),
			}
		}

		tasks = append(tasks, task)
	}

	return tasks
}
