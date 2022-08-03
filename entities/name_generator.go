package entities

import (
	"fmt"
)

type NameGenerator struct {
	prefix  string
	names   map[string]int
	counter int
}

func NewNameGenerator(prefix string) *NameGenerator {
	return &NameGenerator{
		prefix: prefix,
		names:  map[string]int{},
	}
}

func (generator *NameGenerator) GenerateName(suggestedName string) string {
	generatedName := suggestedName
	if generatedName == "" {
		generatedName = fmt.Sprintf("%s #%d", generator.prefix, generator.counter)
		generator.counter++
	}

	generator.names[generatedName]++
	if nameCount := generator.names[generatedName]; nameCount > 1 {
		generatedName = fmt.Sprintf("%s (%d)", generatedName, nameCount)
	}

	return generatedName
}
