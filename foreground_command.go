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

	configDirectory := filepath.Dir(command.ConfigPath)
	tasks := config.PrepareTasks("Task", configDirectory)
	if len(tasks) == 0 {
		const message = "the config does not contain at least one task with phrases"
		return errors.New(message)
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
