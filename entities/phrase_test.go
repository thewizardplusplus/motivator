package entities

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhrase_ExpandText(test *testing.T) {
	type fields struct {
		Text string
	}
	type args struct {
		variables map[string]string
	}

	for _, data := range []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "none of the variables are set",
			fields: fields{
				Text: "variable-one: ${VARIABLE_ONE}; variable-two: ${VARIABLE_TWO}",
			},
			args: args{
				variables: nil,
			},
			want: "variable-one: ; variable-two: ",
		},
		{
			name: "some of the variables are set",
			fields: fields{
				Text: "variable-one: ${VARIABLE_ONE}; variable-two: ${VARIABLE_TWO}",
			},
			args: args{
				variables: map[string]string{"VARIABLE_ONE": "value-one"},
			},
			want: "variable-one: value-one; variable-two: ",
		},
		{
			name: "all of the variables are set",
			fields: fields{
				Text: "variable-one: ${VARIABLE_ONE}; variable-two: ${VARIABLE_TWO}",
			},
			args: args{
				variables: map[string]string{
					"VARIABLE_ONE": "value-one",
					"VARIABLE_TWO": "value-two",
				},
			},
			want: "variable-one: value-one; variable-two: value-two",
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			phrase := Phrase{
				Text: data.fields.Text,
			}
			got := phrase.ExpandText(data.args.variables)

			assert.Equal(test, data.want, got)
		})
	}
}

func TestPhrase_SpinText(test *testing.T) {
	type fields struct {
		Text string
	}

	for _, data := range []struct {
		name    string
		fields  fields
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				Text: "test {one|two|three}",
			},
			want:    "test three",
			wantErr: assert.NoError,
		},
		{
			name: "error",
			fields: fields{
				Text: "test one|two|three}",
			},
			want:    "",
			wantErr: assert.Error,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			rand.Seed(1) // for the reproducibility of the test

			phrase := Phrase{
				Text: data.fields.Text,
			}
			got, err := phrase.SpinText()

			assert.Equal(test, data.want, got)
			data.wantErr(test, err)
		})
	}
}
