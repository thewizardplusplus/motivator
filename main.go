package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
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

	fmt.Println(phrase)
}
