package config

import (
	"github.com/thewizardplusplus/motivator/entities"
)

type Config struct {
	TitleConfig

	Icon      string
	Tasks     []entities.Task
	Variables map[string]string
}
