package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLines(t *testing.T) {
	tests := []struct {
		name  string
		lines [][]string
		want  []problem
	}{
		{
			name: "parses lines correctly",
			lines: [][]string{
				{"5+5", "10"},
				{"1+1", "2"},
			},
			want: []problem{
				{"5+5", "10"},
				{"1+1", "2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLines(tt.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	question string
	answer   string
}
