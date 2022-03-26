package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/m1/gospin"
)

type config struct {
	Phrases []string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	configPath := flag.String("config", "config.json", "")
	flag.Parse()

	// read a config
	configBytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		log.Fatal(err)
	}

	// select a random phrase
	phrase := config.Phrases[rand.Intn(len(config.Phrases))]

	// process the Spintax format
	spinner := gospin.New(nil)
	spin, err := spinner.Spin(phrase)
	if err != nil {
		log.Fatal(err)
	}

	// show a notification
	if err := beeep.Notify("motivator", spin, ""); err != nil {
		log.Fatal(err)
	}
}
