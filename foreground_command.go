package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
	"github.com/m1/gospin"
)

type task struct {
	Cron    string
	Phrases []string
}

type config []task

type foregroundCommand struct {
	configurableCommand
}

func (command foregroundCommand) Run() error {
	// read a config
	configBytes, err := ioutil.ReadFile(command.ConfigPath)
	if err != nil {
		return fmt.Errorf("unable to read a config: %w", err)
	}

	var loadedConfig config
	if err := json.Unmarshal(configBytes, &loadedConfig); err != nil {
		return fmt.Errorf("unable to unmarshal a config: %w", err)
	}

	var filteredConfig config
	for _, task := range loadedConfig {
		if len(task.Phrases) != 0 {
			filteredConfig = append(filteredConfig, task)
		}
	}
	if len(filteredConfig) == 0 {
		return errors.New("there is not at least one task with phrases in the config")
	}

	// start a cron scheduler
	scheduler := gocron.NewScheduler(time.UTC)
	for index, task := range filteredConfig {
		taskCopy := task
		if _, err := scheduler.CronWithSeconds(task.Cron).Do(func() {
			// select a random phrase
			phrase := taskCopy.Phrases[rand.Intn(len(taskCopy.Phrases))]

			// process the Spintax format
			spinner := gospin.New(nil)
			spin, err := spinner.Spin(phrase)
			if err != nil {
				log.Printf("unable to process the Spintax format: %s", err)
				return
			}

			// show a notification
			if err := beeep.Notify(appName, spin, ""); err != nil {
				log.Printf("unable to show a notification: %s", err)
				return
			}
		}); err != nil {
			log.Printf("unable to start a cron scheduler for task #%d: %s", index, err)
		}
	}

	log.Print(markOfShowingStart)
	scheduler.StartBlocking()

	return nil
}
