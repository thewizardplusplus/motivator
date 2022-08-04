package main

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
	"github.com/thewizardplusplus/motivator/config"
	"github.com/thewizardplusplus/motivator/entities"
	systemutils "github.com/thewizardplusplus/motivator/system-utils"
)

type foregroundCommand struct {
	configurableCommand
}

func (command foregroundCommand) Run() error {
	var config config.Config
	if err := systemutils.UnmarshalJSONFile(
		command.ConfigPath,
		&config,
	); err != nil {
		return fmt.Errorf("unable to load the config: %w", err)
	}

	var tasks []entities.Task
	taskNameGenerator := entities.NewNameGenerator("Task")
	configDirectory := filepath.Dir(command.ConfigPath)
	for taskIndex, task := range config.Tasks {
		if len(task.Phrases) == 0 {
			continue
		}

		task.OriginalName = task.Name
		task.Name = taskNameGenerator.GenerateName(taskIndex, task.Name)

		for phraseIndex, phrase := range task.Phrases {
			iconPath := entities.CoalesceStrings(phrase.Icon, task.Icon, config.Icon)
			if iconPath != "" && !filepath.IsAbs(iconPath) {
				iconPath = filepath.Join(configDirectory, iconPath)
			}

			task.Phrases[phraseIndex] = entities.Phrase{
				Icon: iconPath,
				Text: phrase.ExpandText(config.Variables),
			}
		}

		tasks = append(tasks, task)
	}
	if len(tasks) == 0 {
		return errors.New("there is not at least one task with phrases in the config")
	}

	scheduler := gocron.NewScheduler(time.UTC)
	for _, task := range tasks {
		if _, err := task.PlanJob(scheduler, func(task entities.Task) {
			phrase := task.RandomPhrase()
			spunText, err := phrase.SpinText()
			if err != nil {
				log.Printf("unable to process the Spintax format: %s", err)
				return
			}

			title := config.Title(appName, task)
			if err := beeep.Notify(title, spunText, phrase.Icon); err != nil {
				log.Printf("unable to show the notification: %s", err)
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
