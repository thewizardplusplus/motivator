package entities

import (
	"os"
)

type Phrase struct {
	Icon string
	Text string
}

func (phrase Phrase) ExpandText(variables map[string]string) string {
	return os.Expand(phrase.Text, func(name string) string {
		return variables[name]
	})
}
