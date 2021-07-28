package simplefactory_test

import (
	"reflect"
	"testing"

	simplefactory "github.com/kalifun/design-patterns/simple_factory"
)

func TestNewTheTrafficTools(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want simplefactory.TheTrafficTools
	}{
		{
			name: "c",
			args: args{
				t: "c",
			},
			want: &simplefactory.Car{},
		},
		{
			name: "b",
			args: args{
				t: "b",
			},
			want: &simplefactory.Bicycle{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplefactory.NewTheTrafficTools(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTheTrafficTools() = %v, want %v", got, tt.want)
			}
		})
	}
}
