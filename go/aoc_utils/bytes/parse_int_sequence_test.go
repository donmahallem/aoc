package bytes

import "testing"

func TestParseIntSequence(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		sep      byte
		expected []int
		ok       bool
	}{
		{
			name:     "positive numbers",
			input:    "1,2,3,4",
			sep:      ',',
			expected: []int{1, 2, 3, 4},
			ok:       true,
		},
		{
			name:     "mixed sign numbers",
			input:    "-10,0,25,-3",
			sep:      ',',
			expected: []int{-10, 0, 25, -3},
			ok:       true,
		},
		{
			name:     "custom separator",
			input:    "5|6|7",
			sep:      '|',
			expected: []int{5, 6, 7},
			ok:       true,
		},
		{
			name:  "empty input",
			input: "",
			sep:   ',',
			ok:    true,
		},
		{
			name:  "trailing separator",
			input: "8,9,",
			sep:   ',',
			ok:    false,
		},
		{
			name:  "consecutive separators",
			input: "1,,2",
			sep:   ',',
			ok:    false,
		},
		{
			name:  "invalid character",
			input: "1,a,2",
			sep:   ',',
			ok:    false,
		},
		{
			name:  "dangling minus",
			input: "-,1",
			sep:   ',',
			ok:    false,
		},
		{
			name:     "double minus becomes plus at start",
			input:    "--5,10",
			sep:      ',',
			expected: []int{5, 10},
			ok:       true,
		},
		{
			name:  "double minus after value fails",
			input: "1,--2",
			sep:   ',',
			ok:    false,
		},
		{
			name:  "lone double minus fails",
			input: "--",
			sep:   ',',
			ok:    false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			var out []int
			ok := ParseIntSequence[int]([]byte(testCase.input), testCase.sep, &out)
			if ok != testCase.ok {
				t.Fatalf("ParseIntSequence ok=%t, want %t (out=%v)", ok, testCase.ok, out)
			}
			if !ok {
				if len(out) != 0 {
					t.Fatalf("expected out to be cleared on failure, got %v", out)
				}
				return
			}
			if len(out) != len(testCase.expected) {
				t.Fatalf("len(out)=%d, want %d (%v)", len(out), len(testCase.expected), out)
			}
			for i, v := range testCase.expected {
				if out[i] != v {
					t.Fatalf("out[%d]=%d, want %d (%v)", i, out[i], v, out)
				}
			}
		})
	}
}
