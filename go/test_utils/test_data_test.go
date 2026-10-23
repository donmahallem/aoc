package test_utils

import (
	"encoding/json/jsontext"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseTestData(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		extectedErr string
		checkFn     func(t *testing.T, data TestData)
	}{
		{
			name:  "int result",
			input: `{"23":{"01":[{"name":"a","input":"x","part1":{"result":42}}]}}`,
			checkFn: func(t *testing.T, data TestData) {
				if r := data["23"]["01"][0].Part1.Result; r.Kind != ResultInt || r.Int != 42 {
					t.Fatalf("got %+v", r)
				}
			},
		},
		{
			name:  "int array result",
			input: `{"24":{"17":[{"name":"a","file":"f.txt","part1":{"result":[4,6,0]}}]}}`,
			checkFn: func(t *testing.T, data TestData) {
				if r := data["24"]["17"][0].Part1.Result; r.Kind != ResultInts || !reflect.DeepEqual(r.Ints, []int64{4, 6, 0}) {
					t.Fatalf("got %+v", r)
				}
			},
		},
		{
			name:  "string result",
			input: `{"24":{"14":[{"name":"a","input":"x","part1":{"result":"abc"}}]}}`,
			checkFn: func(t *testing.T, data TestData) {
				if r := data["24"]["14"][0].Part1.Result; r.Kind != ResultString || r.Str != "abc" {
					t.Fatalf("got %+v", r)
				}
			},
		},
		{
			name:  "string array result with type hint",
			input: `{"24":{"23":[{"name":"a","input":"x","part2":{"result":["co","de"],"type":"string"}}]}}`,
			checkFn: func(t *testing.T, data TestData) {
				p := data["24"]["23"][0].Part2
				if p.Result.Kind != ResultStrings || !reflect.DeepEqual(p.Result.Strs, []string{"co", "de"}) || p.Type != "string" {
					t.Fatalf("got %+v", p)
				}
			},
		},
		{
			name:        "mixed array result is rejected",
			input:       `{"23":{"01":[{"name":"a","input":"x","part1":{"result":[1,"a"]}}]}}`,
			extectedErr: "result array must contain only integers or only strings",
		},
		{
			name:        "unknown member is rejected",
			input:       `{"23":{"01":[{"name":"a","input":"x","part1":{"result":1},"bogus":true}]}}`,
			extectedErr: "bogus",
		},
		{
			name:        "missing result is rejected",
			input:       `{"23":{"01":[{"name":"a","input":"x","part1":{}}]}}`,
			extectedErr: "result is required",
		},
		{
			name:        "input and file together are rejected",
			input:       `{"23":{"01":[{"name":"a","input":"x","file":"f.txt","part1":{"result":1}}]}}`,
			extectedErr: "exactly one of input or file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := parseTestData([]byte(tt.input), "test.json")
			if tt.extectedErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.extectedErr) {
					t.Fatalf("err = %v, want containing %q", err, tt.extectedErr)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			tt.checkFn(t, data)
		})
	}
}

func TestResultMarshaling(t *testing.T) {
	testData := []Result{
		{Kind: ResultInt, Int: 42},
		{Kind: ResultInts, Ints: []int64{1, 2, 3}},
		{Kind: ResultString, Str: "abc"},
		{Kind: ResultStrings, Strs: []string{"a", "b"}},
	}
	for _, test := range testData {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			var sb strings.Builder
			enc := jsontext.NewEncoder(&sb)
			if err := test.MarshalJSONTo(enc); err != nil {
				t.Fatal(err)
			}
			raw := sb.String()
			dec := jsontext.NewDecoder(strings.NewReader(raw))
			var r2 Result
			if err := r2.UnmarshalJSONFrom(dec); err != nil {
				t.Fatal(err)
			}
			if r2.Kind != test.Kind {
				t.Fatalf("got %+v, want %+v", r2, test)
			}
		})
	}

}
