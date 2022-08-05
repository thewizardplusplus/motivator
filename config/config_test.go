package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thewizardplusplus/motivator/entities"
)

func TestConfig_PrepareTasks(test *testing.T) {
	type fields struct {
		Icon      string
		Tasks     []entities.Task
		Variables map[string]string
	}
	type args struct {
		generatedNamePrefix string
		basicIconPath       string
	}

	for _, data := range []struct {
		name   string
		fields fields
		args   args
		want   []entities.Task
	}{
		{
			name: "without tasks",
			fields: fields{
				Icon:  "config-icon",
				Tasks: nil,
				Variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			args: args{
				generatedNamePrefix: "generated-name-prefix",
				basicIconPath:       "basic-icon-path",
			},
			want: nil,
		},
		{
			name: "with tasks/regular tasks",
			fields: fields{
				Icon: "config-icon",
				Tasks: []entities.Task{
					{
						Name: "Task #1",
						Icon: "task-icon-1",
						Phrases: []entities.Phrase{
							{Icon: "phrase-icon-1", Text: "Phrase #1"},
							{Icon: "phrase-icon-2", Text: "Phrase #2"},
						},
					},
					{
						Name: "Task #2",
						Icon: "task-icon-2",
						Phrases: []entities.Phrase{
							{Icon: "phrase-icon-3", Text: "Phrase #3"},
							{Icon: "phrase-icon-4", Text: "Phrase #4"},
						},
					},
				},
				Variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			args: args{
				generatedNamePrefix: "generated-name-prefix",
				basicIconPath:       "basic-icon-path",
			},
			want: []entities.Task{
				{
					Name:         "Task #1",
					OriginalName: "Task #1",
					Icon:         "task-icon-1",
					Phrases: []entities.Phrase{
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
					Name:         "Task #2",
					OriginalName: "Task #2",
					Icon:         "task-icon-2",
					Phrases: []entities.Phrase{
						{
							Icon: filepath.Join("basic-icon-path", "phrase-icon-3"),
							Text: "Phrase #3",
						},
						{
							Icon: filepath.Join("basic-icon-path", "phrase-icon-4"),
							Text: "Phrase #4",
						},
					},
				},
			},
		},
		{
			name: "with tasks/tasks without phrases",
			fields: fields{
				Icon: "config-icon",
				Tasks: []entities.Task{
					{
						Name: "Task #1",
						Icon: "task-icon-1",
						Phrases: []entities.Phrase{
							{Icon: "phrase-icon-1", Text: "Phrase #1"},
							{Icon: "phrase-icon-2", Text: "Phrase #2"},
						},
					},
					{
						Name:    "Task #2",
						Icon:    "task-icon-2",
						Phrases: nil,
					},
				},
				Variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			args: args{
				generatedNamePrefix: "generated-name-prefix",
				basicIconPath:       "basic-icon-path",
			},
			want: []entities.Task{
				{
					Name:         "Task #1",
					OriginalName: "Task #1",
					Icon:         "task-icon-1",
					Phrases: []entities.Phrase{
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
			},
		},
		{
			name: "with tasks/with generation of task names",
			fields: fields{
				Icon: "config-icon",
				Tasks: []entities.Task{
					{
						Icon: "task-icon-1",
						Phrases: []entities.Phrase{
							{Icon: "phrase-icon-1", Text: "Phrase #1"},
							{Icon: "phrase-icon-2", Text: "Phrase #2"},
						},
					},
					{
						Icon: "task-icon-2",
						Phrases: []entities.Phrase{
							{Icon: "phrase-icon-3", Text: "Phrase #3"},
							{Icon: "phrase-icon-4", Text: "Phrase #4"},
						},
					},
				},
				Variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			args: args{
				generatedNamePrefix: "generated-name-prefix",
				basicIconPath:       "basic-icon-path",
			},
			want: []entities.Task{
				{
					Name: "generated-name-prefix #0",
					Icon: "task-icon-1",
					Phrases: []entities.Phrase{
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
					Name: "generated-name-prefix #1",
					Icon: "task-icon-2",
					Phrases: []entities.Phrase{
						{
							Icon: filepath.Join("basic-icon-path", "phrase-icon-3"),
							Text: "Phrase #3",
						},
						{
							Icon: filepath.Join("basic-icon-path", "phrase-icon-4"),
							Text: "Phrase #4",
						},
					},
				},
			},
		},
		{
			name: "with tasks/the config icon",
			fields: fields{
				Icon: "config-icon",
				Tasks: []entities.Task{
					{
						Name:    "Task #1",
						Phrases: []entities.Phrase{{Text: "Phrase #1"}, {Text: "Phrase #2"}},
					},
					{
						Name:    "Task #2",
						Phrases: []entities.Phrase{{Text: "Phrase #3"}, {Text: "Phrase #4"}},
					},
				},
				Variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			args: args{
				generatedNamePrefix: "generated-name-prefix",
				basicIconPath:       "basic-icon-path",
			},
			want: []entities.Task{
				{
					Name:         "Task #1",
					OriginalName: "Task #1",
					Phrases: []entities.Phrase{
						{
							Icon: filepath.Join("basic-icon-path", "config-icon"),
							Text: "Phrase #1",
						},
						{
							Icon: filepath.Join("basic-icon-path", "config-icon"),
							Text: "Phrase #2",
						},
					},
				},
				{
					Name:         "Task #2",
					OriginalName: "Task #2",
					Phrases: []entities.Phrase{
						{
							Icon: filepath.Join("basic-icon-path", "config-icon"),
							Text: "Phrase #3",
						},
						{
							Icon: filepath.Join("basic-icon-path", "config-icon"),
							Text: "Phrase #4",
						},
					},
				},
			},
		},
		{
			name: "with tasks/with an expansion of phrase text",
			fields: fields{
				Icon: "config-icon",
				Tasks: []entities.Task{
					{
						Name: "Task #1",
						Icon: "task-icon-1",
						Phrases: []entities.Phrase{
							{
								Icon: "phrase-icon-1",
								Text: "Phrase #1 (variable-one: ${VARIABLE_ONE})",
							},
							{
								Icon: "phrase-icon-2",
								Text: "Phrase #2 (variable-two: ${VARIABLE_TWO})",
							},
						},
					},
					{
						Name: "Task #2",
						Icon: "task-icon-2",
						Phrases: []entities.Phrase{
							{
								Icon: "phrase-icon-3",
								Text: "Phrase #3 (variable-one: ${VARIABLE_ONE})",
							},
							{
								Icon: "phrase-icon-4",
								Text: "Phrase #4 (variable-two: ${VARIABLE_TWO})",
							},
						},
					},
				},
				Variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			args: args{
				generatedNamePrefix: "generated-name-prefix",
				basicIconPath:       "basic-icon-path",
			},
			want: []entities.Task{
				{
					Name:         "Task #1",
					OriginalName: "Task #1",
					Icon:         "task-icon-1",
					Phrases: []entities.Phrase{
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
				{
					Name:         "Task #2",
					OriginalName: "Task #2",
					Icon:         "task-icon-2",
					Phrases: []entities.Phrase{
						{
							Icon: filepath.Join("basic-icon-path", "phrase-icon-3"),
							Text: "Phrase #3 (variable-one: value-one)",
						},
						{
							Icon: filepath.Join("basic-icon-path", "phrase-icon-4"),
							Text: "Phrase #4 (variable-two: value-two)",
						},
					},
				},
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			config := Config{
				Icon:      data.fields.Icon,
				Tasks:     data.fields.Tasks,
				Variables: data.fields.Variables,
			}
			got := config.PrepareTasks(
				data.args.generatedNamePrefix,
				data.args.basicIconPath,
			)

			assert.Equal(test, data.want, got)
		})
	}
}
