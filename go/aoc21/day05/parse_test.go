package day05

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input   string
		want    []coords
		wantErr bool
	}{
		{"0,0 -> 1,1\n", []coords{{0, 0, 1, 1, DirectionDiagonal}}, false},
		{"0,1 -> 1,5\n", []coords{{0, 1, 1, 5, DirectionOther}}, false},
		{"0,0 -> 0,1\n", []coords{{0, 0, 0, 1, DirectionVertical}}, false},
		{"123,1234 -> 1234,1234\n", []coords{{123, 1234, 1234, 1234, DirectionHorizontal}}, false},
		{"98765,1234 -> 1234,2\n", []coords{{1234, 1234, 98765, 2, DirectionOther}}, false},
	}
	for _, tt := range tests {
		got, err := parseInput(strings.NewReader(tt.input))
		if (err != nil) != tt.wantErr {
			t.Errorf("parseInput(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("parseInput(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
