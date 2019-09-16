package apufferi

import (
	"reflect"
	"testing"
)

var interfaceStringArray = []string{
	"12345",
	"true",
	"test3",
}

var interfaceObjectArray = []interface{}{
	12345,
	true,
	"test3",
}

var interfaceInvalid = map[string]string{
	"invalid": "input",
}

func TestToStringArray(t *testing.T) {
	type args struct {
		element interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test null input",
			args: args{
				element: nil,
			},
			want: nil,
		},
		{
			name: "Test empty input",
			args: args{
				element: make([]string, 0),
			},
			want: make([]string, 0),
		},
		{
			name: "Test all valid input",
			args: args{
				element: interfaceStringArray,
			},
			want: interfaceStringArray,
		},
		{
			name: "Test mixed input",
			args: args{
				element: interfaceObjectArray,
			},
			want: interfaceStringArray,
		},
		{
			name: "Test single string input",
			args: args{
				element: "test",
			},
			want: []string{"test"},
		},
		{
			name: "Test invalid type",
			args: args{
				element: interfaceInvalid,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStringArray(tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
