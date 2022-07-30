package entities

import (
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
