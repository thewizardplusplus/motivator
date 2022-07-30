package main

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
	"github.com/thewizardplusplus/motivator/entities"
	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type config struct {
	Icon                string
	Tasks               []entities.Task
	Variables           map[string]string
	HideAppName         bool
	UseOriginalTaskName bool
}

type foregroundCommand struct {
	configurableCommand
}

func (command foregroundCommand) Run() error {
	var config config
	if err := systemutils.UnmarshalJSONFile(
		command.ConfigPath,
		&config,
	); err != nil {
		return fmt.Errorf("unable to load the config: %w", err)
	}

	var tasks []entities.Task
	taskNames := make(map[string]int)
	configDirectory := filepath.Dir(command.ConfigPath)
	for taskIndex, task := range config.Tasks {
		if len(task.Phrases) == 0 {
			continue
		}

		task.OriginalName = task.Name

		if task.Name == "" {
			task.Name = fmt.Sprintf("Task #%d", taskIndex+1)
		}

		taskNames[task.Name]++
		if count := taskNames[task.Name]; count > 1 {
			task.Name = fmt.Sprintf("%s #%d", task.Name, count)
		}

		alternativeIconPath := task.Icon
		if alternativeIconPath == "" {
			alternativeIconPath = config.Icon
		}
		for phraseIndex, phrase := range task.Phrases {
			iconPath := phrase.Icon
			if iconPath == "" {
				iconPath = alternativeIconPath
			}
			if iconPath != "" && !filepath.IsAbs(iconPath) {
				iconPath = filepath.Join(configDirectory, iconPath)
			}
			task.Phrases[phraseIndex].Icon = iconPath

			task.Phrases[phraseIndex].Text = phrase.ExpandText(config.Variables)
		}

		tasks = append(tasks, task)
	}
	if len(tasks) == 0 {
		return errors.New("there is not at least one task with phrases in the config")
	}

	// start a cron scheduler
	scheduler := gocron.NewScheduler(time.UTC)
	for _, task := range tasks {
		if _, err := task.PlanJob(scheduler, func(task entities.Task) {
			// select a random phrase
			phrase := task.RandomPhrase()

			// process the Spintax format
			spunText, err := phrase.SpinText()
			if err != nil {
				log.Printf("unable to process the Spintax format: %s", err)
				return
			}

			// show a notification
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
			title := strings.Join(titleParts, " | ")
			if err := beeep.Notify(title, spunText, phrase.Icon); err != nil {
				log.Printf("unable to show a notification: %s", err)
				return
			}
		}); err != nil {
			const message = "unable to start the job scheduler for task %q: %s"
			log.Printf(message, task.Name, err)
		}
	}
	if len(scheduler.Jobs()) == 0 {
		return errors.New("unable to start the job scheduler for at least one task")
	}

	log.Print(markOfShowingStart)
	scheduler.StartBlocking()

	return nil
}
