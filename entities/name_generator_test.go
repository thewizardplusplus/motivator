package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNameGenerator(test *testing.T) {
	generator := NewNameGenerator("test")

	assert.Equal(test, "test", generator.prefix)
	assert.Equal(test, map[string]int{}, generator.names)
	assert.Equal(test, 0, generator.counter)
}
