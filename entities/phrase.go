package entities

import (
	"fmt"
	"os"

	"github.com/m1/gospin"
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

func (phrase Phrase) SpinText() (string, error) {
	spunText, err :=
		gospin.New(&gospin.Config{UseGlobalRand: true}).Spin(phrase.Text)
	if err != nil {
		return "", fmt.Errorf("unable to process the Spintax format: %w", err)
	}

	return spunText, nil
}
