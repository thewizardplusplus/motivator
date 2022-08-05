package entities

import (
	"context"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestTask_PreparePhrases(test *testing.T) {
	type fields struct {
		Icon    string
		Phrases []Phrase
	}
	type args struct {
		defaultIcon   string
		basicIconPath string
		variables     map[string]string
	}

	for _, data := range []struct {
		name   string
		fields fields
		args   args
		want   []Phrase
	}{
		{
			name: "without phrases",
			fields: fields{
				Phrases: nil,
			},
			args: args{
				defaultIcon:   "default-icon",
				basicIconPath: "basic-icon-path",
				variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			want: nil,
		},
		{
			name: "with phrases/phrase icons",
			fields: fields{
				Icon: "task-icon",
				Phrases: []Phrase{
					{Icon: "phrase-icon-1", Text: "Phrase #1"},
					{Icon: "phrase-icon-2", Text: "Phrase #2"},
				},
			},
			args: args{
				defaultIcon:   "default-icon",
				basicIconPath: "basic-icon-path",
				variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			want: []Phrase{
				{
					Icon: filepath.Join("basic-icon-path", "phrase-icon-1"),
					Text: "Phrase #1",
				},
				{
					Icon: filepath.Join("basic-icon-path", "phrase-icon-2"),
					Text: "Phrase #2",
				},
			},
		},
		{
			name: "with phrases/task icons",
			fields: fields{
				Icon:    "task-icon",
				Phrases: []Phrase{{Text: "Phrase #1"}, {Text: "Phrase #2"}},
			},
			args: args{
				defaultIcon:   "default-icon",
				basicIconPath: "basic-icon-path",
				variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			want: []Phrase{
				{Icon: filepath.Join("basic-icon-path", "task-icon"), Text: "Phrase #1"},
				{Icon: filepath.Join("basic-icon-path", "task-icon"), Text: "Phrase #2"},
			},
		},
		{
			name: "with phrases/default icons",
			fields: fields{
				Phrases: []Phrase{{Text: "Phrase #1"}, {Text: "Phrase #2"}},
			},
			args: args{
				defaultIcon:   "default-icon",
				basicIconPath: "basic-icon-path",
				variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			want: []Phrase{
				{Icon: filepath.Join("basic-icon-path", "default-icon"), Text: "Phrase #1"},
				{Icon: filepath.Join("basic-icon-path", "default-icon"), Text: "Phrase #2"},
			},
		},
		{
			name: "with phrases/without icons",
			fields: fields{
				Phrases: []Phrase{{Text: "Phrase #1"}, {Text: "Phrase #2"}},
			},
			args: args{
				defaultIcon:   "",
				basicIconPath: "basic-icon-path",
				variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			want: []Phrase{{Text: "Phrase #1"}, {Text: "Phrase #2"}},
		},
		{
			name: "with phrases/absolute paths to icons",
			fields: fields{
				Icon: "task-icon",
				Phrases: []Phrase{
					{Icon: filepath.Join(os.TempDir(), "phrase-icon-1"), Text: "Phrase #1"},
					{Icon: filepath.Join(os.TempDir(), "phrase-icon-2"), Text: "Phrase #2"},
				},
			},
			args: args{
				defaultIcon:   "default-icon",
				basicIconPath: "basic-icon-path",
				variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			want: []Phrase{
				{Icon: filepath.Join(os.TempDir(), "phrase-icon-1"), Text: "Phrase #1"},
				{Icon: filepath.Join(os.TempDir(), "phrase-icon-2"), Text: "Phrase #2"},
			},
		},
		{
			name: "with phrases/with an expansion of phrase text",
			fields: fields{
				Icon: "task-icon",
				Phrases: []Phrase{
					{Icon: "phrase-icon-1", Text: "Phrase #1 (variable-one: ${VARIABLE_ONE})"},
					{Icon: "phrase-icon-2", Text: "Phrase #2 (variable-two: ${VARIABLE_TWO})"},
				},
			},
			args: args{
				defaultIcon:   "default-icon",
				basicIconPath: "basic-icon-path",
				variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			want: []Phrase{
				{
					Icon: filepath.Join("basic-icon-path", "phrase-icon-1"),
					Text: "Phrase #1 (variable-one: value-one)",
				},
				{
					Icon: filepath.Join("basic-icon-path", "phrase-icon-2"),
					Text: "Phrase #2 (variable-two: value-two)",
				},
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			task := Task{
				Icon:    data.fields.Icon,
				Phrases: data.fields.Phrases,
			}
			got := task.PreparePhrases(
				data.args.defaultIcon,
				data.args.basicIconPath,
				data.args.variables,
			)

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
		taskHandler TaskHandler
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
				taskHandler: &MockTaskHandler{},
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
				taskHandler: &MockTaskHandler{},
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
				taskHandler: &MockTaskHandler{},
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
				taskHandler: &MockTaskHandler{},
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
				taskHandler: &MockTaskHandler{},
			},
			wantNextRun: time.Time{},
			wantErr:     assert.Error,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			// use the `wait for schedule` mode so that jobs are not started immediately
			scheduler := gocron.NewScheduler(time.UTC).WaitForSchedule()
			defer scheduler.Stop()

			task := Task{
				Cron:  data.fields.Cron,
				Delay: data.fields.Delay,
			}
			job, err := task.PlanJob(scheduler, data.args.taskHandler.HandleTask)

			scheduler.StartAsync()

			mock.AssertExpectationsForObjects(test, data.args.taskHandler)
			if job != nil {
				assert.WithinDuration(test, data.wantNextRun, job.NextRun(), time.Minute)
			}
			data.wantErr(test, err)
		})
	}
}

func TestTask_PlanJob_withRunningJob(test *testing.T) {
	task := Task{
		Delay: "1h",
	}

	taskHandler := &MockTaskHandler{}
	taskHandler.On("HandleTask", task).Return()

	// use the `start immediately` mode so that jobs are started immediately
	scheduler := gocron.NewScheduler(time.UTC).StartImmediately()
	defer scheduler.Stop()

	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second)
	job, err := task.PlanJob(scheduler, func(task Task) {
		defer ctxCancel()

		taskHandler.HandleTask(task)
	})

	scheduler.StartAsync()
	<-ctx.Done()

	mock.AssertExpectationsForObjects(test, taskHandler)
	if job != nil {
		wantNextRun := time.Now().Add(time.Hour)
		assert.WithinDuration(test, wantNextRun, job.NextRun(), time.Minute)
	}
	assert.NoError(test, err)
}
