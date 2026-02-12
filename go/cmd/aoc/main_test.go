package main

import (
	"testing"
)

func TestParsePartSelector(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantYear int
		wantDay  int
		wantPart int
		wantErr  bool
	}{
		{
			name:     "Valid args",
			args:     []string{"24", "1", "1"},
			wantYear: 24,
			wantDay:  1,
			wantPart: 1,
			wantErr:  false,
		},
		{
			name:     "Full Year 2024",
			args:     []string{"2024", "1", "1"},
			wantYear: 24,
			wantDay:  1,
			wantPart: 1,
			wantErr:  false,
		},
		{
			name:    "Not enough args",
			args:    []string{"2024", "1"},
			wantErr: true,
		},
		{
			name:    "Invalid year",
			args:    []string{"abcd", "1", "1"},
			wantErr: true,
		},
		{
			name:    "Invalid day",
			args:    []string{"2024", "foo", "1"},
			wantErr: true,
		},
		{
			name:    "Invalid part",
			args:    []string{"2024", "1", "bar"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y, d, p, err := parsePartSelector(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePartSelector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if y != tt.wantYear {
					t.Errorf("parsePartSelector() Year = %v, want %v", y, tt.wantYear)
				}
				if d != tt.wantDay {
					t.Errorf("parsePartSelector() Day = %v, want %v", d, tt.wantDay)
				}
				if p != tt.wantPart {
					t.Errorf("parsePartSelector() Part = %v, want %v", p, tt.wantPart)
				}
			}
		})
	}
}
