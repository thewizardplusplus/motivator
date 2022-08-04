package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNameGenerator(test *testing.T) {
	generator := NewNameGenerator("test")

	assert.Equal(test, "test", generator.prefix)
	assert.Equal(test, map[string]int{}, generator.names)
}

func TestNameGenerator_GenerateName(test *testing.T) {
	type fields struct {
		prefix string
		names  map[string]int
	}
	type args struct {
		nameIndex     int
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
				prefix: "test",
				names:  map[string]int{},
			},
			args: args{
				nameIndex:     23,
				suggestedName: "name",
			},
			want: "name",
			wantGenerator: &NameGenerator{
				prefix: "test",
				names:  map[string]int{"name": 1},
			},
		},
		{
			name: "empty name",
			fields: fields{
				prefix: "test",
				names:  map[string]int{},
			},
			args: args{
				nameIndex:     23,
				suggestedName: "",
			},
			want: "test #23",
			wantGenerator: &NameGenerator{
				prefix: "test",
				names:  map[string]int{"test #23": 1},
			},
		},
		{
			name: "duplicated name",
			fields: fields{
				prefix: "test",
				names:  map[string]int{"name": 2},
			},
			args: args{
				nameIndex:     23,
				suggestedName: "name",
			},
			want: "name (3)",
			wantGenerator: &NameGenerator{
				prefix: "test",
				names:  map[string]int{"name": 3},
			},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			generator := &NameGenerator{
				prefix: data.fields.prefix,
				names:  data.fields.names,
			}
			got := generator.GenerateName(data.args.nameIndex, data.args.suggestedName)

			assert.Equal(test, data.want, got)
			assert.Equal(test, data.wantGenerator, generator)
		})
	}
}
