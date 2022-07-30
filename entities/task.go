package entities

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/go-co-op/gocron"
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

func (task Task) SelectedName() string {
	name := task.Name
	if task.UseOriginalName {
		name = task.OriginalName
	}

	return name
}

func (task Task) RandomPhrase() Phrase {
	return task.Phrases[rand.Intn(len(task.Phrases))]
}

func (task Task) PlanJob(
	scheduler *gocron.Scheduler,
	taskHandler func(task Task),
) (*gocron.Job, error) {
	type jobPlannerType func(jobHandler interface{}, parameters ...interface{}) (
		*gocron.Job,
		error,
	)
	type cronParserType func(cronExpression string) *gocron.Scheduler

	var jobPlanner jobPlannerType
	if task.Cron != "" {
		var cronParser cronParserType
		const cronWithSecondsFieldCount = 6
		if fields := strings.Fields(task.Cron); len(fields) == cronWithSecondsFieldCount { // nolint: lll
			cronParser = scheduler.CronWithSeconds
		} else {
			cronParser = scheduler.Cron
		}

		jobPlanner = cronParser(task.Cron).Do
	} else {
		jobPlanner = scheduler.Every(task.Delay).Do
	}

	job, err := jobPlanner(func() { taskHandler(task) })
	if err != nil {
		return nil, fmt.Errorf("unable to start the job scheduler: %w", err)
	}

	return job, nil
}
