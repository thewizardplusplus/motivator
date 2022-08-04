package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoalesceStrings(test *testing.T) {
	type args struct {
		strings []string
	}

	for _, data := range []struct {
		name string
		args args
		want string
	}{
		{
			name: "without strings",
			args: args{
				strings: nil,
			},
			want: "",
		},
		{
			name: "with non-empty strings",
			args: args{
				strings: []string{"one", "two", "three"},
			},
			want: "one",
		},
		{
			name: "with mixed strings",
			args: args{
				strings: []string{"", "two", "three"},
			},
			want: "two",
		},
		{
			name: "with empty strings",
			args: args{
				strings: []string{"", "", ""},
			},
			want: "",
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			got := CoalesceStrings(data.args.strings...)

			assert.Equal(test, data.want, got)
		})
	}
}
