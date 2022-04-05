package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
	"github.com/m1/gospin"
)

type phrase struct {
	Icon string
	Text string
}

type task struct {
	Name    string
	Icon    string
	Cron    string
	Phrases []phrase
}

type config struct {
	Icon      string
	Tasks     []task
	Variables map[string]string
}

type foregroundCommand struct {
	configurableCommand
}

func (command foregroundCommand) Run() error {
	// read a config
	configBytes, err := ioutil.ReadFile(command.ConfigPath)
	if err != nil {
		return fmt.Errorf("unable to read a config: %w", err)
	}

	var config config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		return fmt.Errorf("unable to unmarshal a config: %w", err)
	}

	var tasks []task
	taskNames := make(map[string]int)
	configDirectory := filepath.Dir(command.ConfigPath)
	for taskIndex, task := range config.Tasks {
		if len(task.Phrases) == 0 {
			continue
		}

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

			task.Phrases[phraseIndex].Text =
				os.Expand(phrase.Text, func(name string) string {
					return config.Variables[name]
				})
		}

		tasks = append(tasks, task)
	}
	if len(tasks) == 0 {
		return errors.New("there is not at least one task with phrases in the config")
	}

	// start a cron scheduler
	scheduler := gocron.NewScheduler(time.UTC)
	for _, task := range tasks {
		taskCopy := task
		if _, err := scheduler.CronWithSeconds(task.Cron).Do(func() {
			// select a random phrase
			phrase := taskCopy.Phrases[rand.Intn(len(taskCopy.Phrases))]

			// process the Spintax format
			spinner := gospin.New(nil)
			spin, err := spinner.Spin(phrase.Text)
			if err != nil {
				log.Printf("unable to process the Spintax format: %s", err)
				return
			}

			// show a notification
			title := fmt.Sprintf("%s | %s", appName, taskCopy.Name)
			if err := beeep.Notify(title, spin, phrase.Icon); err != nil {
				log.Printf("unable to show a notification: %s", err)
				return
			}
		}); err != nil {
			log.Printf(
				"unable to start a cron scheduler for task %q: %s",
				taskCopy.Name,
				err,
			)
		}
	}
	if len(scheduler.Jobs()) == 0 {
		return errors.New("unable to start a cron scheduler for at least one task")
	}

	log.Print(markOfShowingStart)
	scheduler.StartBlocking()

	return nil
}
