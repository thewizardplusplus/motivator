package entities

import (
	"math/rand"
	"testing"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/stretchr/testify/assert"
)

func TestTask_SelectedName(test *testing.T) {
	type fields struct {
		Name            string
		OriginalName    string
		UseOriginalName bool
	}

	for _, data := range []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: `use the "Name" field`,
			fields: fields{
				Name:            "Name",
				OriginalName:    "OriginalName",
				UseOriginalName: false,
			},
			want: "Name",
		},
		{
			name: `use the "OriginalName" field`,
			fields: fields{
				Name:            "Name",
				OriginalName:    "OriginalName",
				UseOriginalName: true,
			},
			want: "OriginalName",
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			task := Task{
				Name:            data.fields.Name,
				OriginalName:    data.fields.OriginalName,
				UseOriginalName: data.fields.UseOriginalName,
			}
			got := task.SelectedName()

			assert.Equal(test, data.want, got)
		})
	}
}

func TestTask_RandomPhrase(test *testing.T) {
	rand.Seed(1) // for the reproducibility of the test

	task := Task{
		Phrases: []Phrase{{Text: "one"}, {Text: "two"}, {Text: "three"}},
	}
	got := task.RandomPhrase()

	assert.Equal(test, Phrase{Text: "three"}, got)
}

func TestTask_PlanJob(test *testing.T) {
	type fields struct {
		Cron  string
		Delay string
	}
	type args struct {
		taskHandler func(task Task)
	}

	for _, data := range []struct {
		name        string
		fields      fields
		args        args
		wantNextRun time.Time
		wantErr     assert.ErrorAssertionFunc
	}{
		{
			name: "success/cron with seconds",
			fields: fields{
				Cron: "0 0 * * * *",
			},
			args: args{
				taskHandler: func(task Task) {},
			},
			wantNextRun: time.Now().Truncate(time.Hour).Add(time.Hour),
			wantErr:     assert.NoError,
		},
		{
			name: "success/cron without seconds",
			fields: fields{
				Cron: "0 * * * *",
			},
			args: args{
				taskHandler: func(task Task) {},
			},
			wantNextRun: time.Now().Truncate(time.Hour).Add(time.Hour),
			wantErr:     assert.NoError,
		},
		{
			name: "success/delay",
			fields: fields{
				Delay: "1h",
			},
			args: args{
				taskHandler: func(task Task) {},
			},
			wantNextRun: time.Now().Add(time.Hour),
			wantErr:     assert.NoError,
		},
		{
			name: "error/cron",
			fields: fields{
				Cron: "incorrect",
			},
			args: args{
				taskHandler: func(task Task) {},
			},
			wantNextRun: time.Time{},
			wantErr:     assert.Error,
		},
		{
			name: "error/delay",
			fields: fields{
				Delay: "incorrect",
			},
			args: args{
				taskHandler: func(task Task) {},
			},
			wantNextRun: time.Time{},
			wantErr:     assert.Error,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			scheduler := gocron.NewScheduler(time.UTC)
			defer scheduler.Stop()

			task := Task{
				Cron:  data.fields.Cron,
				Delay: data.fields.Delay,
			}
			job, err := task.PlanJob(scheduler, data.args.taskHandler)

			scheduler.StartAsync()

			if job != nil {
				assert.WithinDuration(test, data.wantNextRun, job.NextRun(), time.Minute)
			}
			data.wantErr(test, err)
		})
	}
}
