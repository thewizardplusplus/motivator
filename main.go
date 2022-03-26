package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type config struct {
	Phrases []string
}

func main() {
	configPath := flag.String("config", "config.json", "")
	flag.Parse()

	configBytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}
