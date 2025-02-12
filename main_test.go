// main_test.go
package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		want  map[string]int
		name  string
		input string
	}{
		{
			name:  "basic counting",
			input: "hello world hello",
			want:  map[string]int{"hello": 2, "world": 1},
		},
		{
			name:  "empty string",
			input: "",
			want:  map[string]int{},
		},
		{
			name:  "mixed case",
			input: "Hello WORLD hello",
			want:  map[string]int{"hello": 2, "world": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countWords(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
