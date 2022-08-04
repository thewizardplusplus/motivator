package config

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/thewizardplusplus/motivator/entities"
	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type Config struct {
	TitleConfig

	Icon      string
	Tasks     []entities.Task
	Variables map[string]string
}

func LoadConfig(configPath string, generatedNamePrefix string) (Config, error) {
	var config Config
	if err := systemutils.UnmarshalJSONFile(configPath, &config); err != nil {
		return Config{}, fmt.Errorf("unable to load the config: %w", err)
	}

	basicIconPath := filepath.Dir(configPath)
	config.Tasks = config.PrepareTasks(generatedNamePrefix, basicIconPath)
	if len(config.Tasks) == 0 {
		const message = "the config does not contain at least one task with phrases"
		return Config{}, errors.New(message)
	}

	return config, nil
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
