package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thewizardplusplus/motivator/entities"
)

func TestTitleConfig_Title(test *testing.T) {
	type fields struct {
		HideAppName         bool
		UseOriginalTaskName bool
	}
	type args struct {
		appName string
		task    entities.Task
	}

	for _, data := range []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "a regular case",
			fields: fields{
				HideAppName:         false,
				UseOriginalTaskName: false,
			},
			args: args{
				appName: "app-name",
				task: entities.Task{
					Name:            "task-name",
					OriginalName:    "task-original-name",
					UseOriginalName: false,
				},
			},
			want: "app-name | task-name",
		},
		{
			name: "use an original task name/" +
				"via the `entities.Task.UseOriginalName` field",
			fields: fields{
				HideAppName:         false,
				UseOriginalTaskName: false,
			},
			args: args{
				appName: "app-name",
				task: entities.Task{
					Name:            "task-name",
					OriginalName:    "task-original-name",
					UseOriginalName: true,
				},
			},
			want: "app-name | task-original-name",
		},
		{
			name: "use an original task name/" +
				"via the `config.TitleConfig.UseOriginalTaskName` field",
			fields: fields{
				HideAppName:         false,
				UseOriginalTaskName: true,
			},
			args: args{
				appName: "app-name",
				task: entities.Task{
					Name:            "task-name",
					OriginalName:    "task-original-name",
					UseOriginalName: false,
				},
			},
			want: "app-name | task-original-name",
		},
		{
			name: "without an app name",
			fields: fields{
				HideAppName:         true,
				UseOriginalTaskName: false,
			},
			args: args{
				appName: "app-name",
				task: entities.Task{
					Name:            "task-name",
					OriginalName:    "task-original-name",
					UseOriginalName: false,
				},
			},
			want: "task-name",
		},
		{
			name: "without a task name",
			fields: fields{
				HideAppName:         false,
				UseOriginalTaskName: false,
			},
			args: args{
				appName: "app-name",
				task: entities.Task{
					Name:            "",
					OriginalName:    "task-original-name",
					UseOriginalName: false,
				},
			},
			want: "app-name",
		},
		{
			name: "without any names",
			fields: fields{
				HideAppName:         false,
				UseOriginalTaskName: false,
			},
			args: args{
				appName: "",
				task: entities.Task{
					Name:            "",
					OriginalName:    "task-original-name",
					UseOriginalName: false,
				},
			},
			want: "",
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			config := TitleConfig{
				HideAppName:         data.fields.HideAppName,
				UseOriginalTaskName: data.fields.UseOriginalTaskName,
			}
			got := config.Title(data.args.appName, data.args.task)

			assert.Equal(test, data.want, got)
		})
	}
}
