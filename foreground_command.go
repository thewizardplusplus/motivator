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

type config struct {
	Cron    string
	Phrases []string
}

type foregroundCommand struct {
	ConfigPath string `kong:"required,short='c',name='config',default='config.json',help='Config path.'"` // nolint: lll
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
	if len(config.Phrases) == 0 {
		return errors.New("config has no phrases")
	}

	// start a cron scheduler
	scheduler := gocron.NewScheduler(time.UTC)
	if _, err := scheduler.CronWithSeconds(config.Cron).Do(func() {
		// select a random phrase
		phrase := config.Phrases[rand.Intn(len(config.Phrases))]

		// process the Spintax format
		spinner := gospin.New(nil)
		spin, err := spinner.Spin(phrase)
		if err != nil {
			log.Printf("unable to process the Spintax format: %s", err)
			return
		}

		// show a notification
		if err := beeep.Notify("motivator", spin, ""); err != nil {
			log.Printf("unable to show a notification: %s", err)
			return
		}
	}); err != nil {
		return fmt.Errorf("unable to start a cron scheduler: %w", err)
	}

	scheduler.StartBlocking()
	return nil
}
