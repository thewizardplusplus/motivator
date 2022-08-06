package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
	cliutils "github.com/thewizardplusplus/motivator/cli-utils"
	"github.com/thewizardplusplus/motivator/config"
	"github.com/thewizardplusplus/motivator/entities"
)

type foregroundCommand struct {
	cliutils.ExecutableInfoCommand `kong:"-"`
	configurableCommand
}

func (command foregroundCommand) Run() error {
	config, err := config.LoadConfig(command.ConfigPath, "Task")
	if err != nil {
		return fmt.Errorf("unable to load the config: %w", err)
	}

	scheduler := gocron.NewScheduler(time.UTC)
	for _, task := range config.Tasks {
		if _, err := task.PlanJob(scheduler, func(task entities.Task) {
			phrase := task.RandomPhrase()
			spunText, err := phrase.SpinText()
			if err != nil {
				log.Printf("unable to process the Spintax format: %s", err)
				return
			}

			title := config.Title(command.ExecutableInfo.Name, task)
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

	log.Print(notificationDisplayStartMark)
	scheduler.StartBlocking()

	return nil
}
