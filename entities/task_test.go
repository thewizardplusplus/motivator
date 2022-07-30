package entities

import (
	"math/rand"
	"testing"

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
