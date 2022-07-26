package entities

import (
	"math/rand"
)

type Task struct {
	Name            string
	OriginalName    string `json:"-"`
	UseOriginalName bool
	Icon            string
	Cron            string
	Delay           string
	Phrases         []Phrase
}

func (task Task) RandomPhrase() Phrase {
	return task.Phrases[rand.Intn(len(task.Phrases))]
}
