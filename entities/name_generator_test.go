package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNameGenerator(test *testing.T) {
	generator := NewNameGenerator("test")

	assert.Equal(test, "test", generator.prefix)
	assert.Equal(test, map[string]int{}, generator.names)
	assert.Equal(test, 0, generator.counter)
}

func TestNameGenerator_GenerateName(test *testing.T) {
	type fields struct {
		prefix  string
		names   map[string]int
		counter int
	}
	type args struct {
		suggestedName string
	}

	for _, data := range []struct {
		name          string
		fields        fields
		args          args
		want          string
		wantGenerator *NameGenerator
	}{
		{
			name: "original name",
			fields: fields{
				prefix:  "test",
				names:   map[string]int{},
				counter: 0,
			},
			args: args{
				suggestedName: "name",
			},
			want: "name",
			wantGenerator: &NameGenerator{
				prefix:  "test",
				names:   map[string]int{"name": 1},
				counter: 0,
			},
		},
		{
			name: "empty name",
			fields: fields{
				prefix:  "test",
				names:   map[string]int{},
				counter: 2,
			},
			args: args{
				suggestedName: "",
			},
			want: "test #2",
			wantGenerator: &NameGenerator{
				prefix:  "test",
				names:   map[string]int{"test #2": 1},
				counter: 3,
			},
		},
		{
			name: "duplicated name",
			fields: fields{
				prefix:  "test",
				names:   map[string]int{"name": 2},
				counter: 0,
			},
			args: args{
				suggestedName: "name",
			},
			want: "name (3)",
			wantGenerator: &NameGenerator{
				prefix:  "test",
				names:   map[string]int{"name": 3},
				counter: 0,
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			generator := &NameGenerator{
				prefix:  data.fields.prefix,
				names:   data.fields.names,
				counter: data.fields.counter,
			}
			got := generator.GenerateName(data.args.suggestedName)

			assert.Equal(test, data.want, got)
			assert.Equal(test, data.wantGenerator, generator)
		})
	}
}
